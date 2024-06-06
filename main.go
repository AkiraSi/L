package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/nats-io/stan.go"
)

func ConnectToPostgres(config map[string]string) (*pgxpool.Pool, error) {
	host := config["host"]
	port := config["port"]
	user := config["user"]
	password := config["password"]
	database := config["database"]
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s database=%s",
		host, port, user, password, database,
	)
	pool, err := pgxpool.Connect(context.Background(), connectionString)
	if err != nil {
		return nil, err
	}
	return pool, nil
}

type Order struct {
	OrderUID          string   `json:"order_uid"`
	TrackNumber       string   `json:"track_number"`
	Entry             string   `json:"entry"`
	Locale            string   `json:"locale"`
	InternalSignature string   `json:"internal_signature"`
	CustomerID        string   `json:"customer_id"`
	Delivery_Service  string   `json:"delivery_service"`
	Shardkey          string   `json:"shardkey"`
	SmID              int      `json:"sm_id"`
	Date_Created      string   `json:"date_created"`
	Payment           Payment  `json:"payment"`
	Items             []Items  `json:"items"`
	Delivery          Delivery `json:"delivery"`
	Oof_shard         string   `json:"oof_shard"`
}

type Delivery struct {
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Zip         string `json:"zip"`
	City        string `json:"city"`
	Address     string `json:"address"`
	Region      string `json:"region"`
	Email       string `json:"email"`
	OrderUID    string `json:"order_uid"`
	TrackNumber string `json:"track_number"`
}

type Payment struct {
	OrderUID     string `json:"order_uid"`
	TrackNumber  string `json:"track_number"`
	RequestID    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       int    `json:"amount"`
	PaymentDt    int    `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost int    `json:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total"`
	CustomFee    int    `json:"custom_fee"`
}

type Items struct {
	ChrtID      int    `json:"chrt_id"`
	OrderUID    string `json:"order_uid"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
	Rid         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  int    `json:"total_price"`
	NmID        int    `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"`
}

var cache = sync.Map{}

func get(oid string) (order Order, i int) {
	if v, ok := cache.Load(oid); ok {
		order = v.(Order)
		return order, 0
	}
	order, i = getdb(oid)
	if i == 0 {
		cache.Store(oid, order)
		config := map[string]string{
			"host":     "localhost",
			"port":     "5432",
			"user":     "postgres",
			"password": "1",
			"database": "L",
		}
		pool, err := ConnectToPostgres(config)
		if err != nil {
			fmt.Println("не успешко")
		}
		defer pool.Close()
		var id int
		err = pool.QueryRow(context.Background(), `SELECT MAX(id) FROM public.cache`).Scan(&id)
		if err != nil {
			log.Printf("Unable to insert data (Order): %vn", err)
		}
		err = pool.QueryRow(context.Background(), `INSERT INTO public.cache(id, order_uid) VALUES ($1, $2);`, strconv.Itoa(id+1), oid).Scan()
		if err != nil {
			errStr := err.Error()
			if errStr == "no rows in result set" {
				fmt.Println("good - cache")
			} else {
				log.Printf("Unable to insert data (Order): %v\n", err)
			}
		}
	}
	return order, i
}

func restoreCache() {
	config := map[string]string{
		"host":     "localhost",
		"port":     "5432",
		"user":     "postgres",
		"password": "1",
		"database": "L",
	}
	pool, err := ConnectToPostgres(config)
	if err != nil {
		fmt.Println("не успешко")
	}
	defer pool.Close()
	rows, err := pool.Query(context.Background(), `SELECT order_uid FROM public.cache ORDER BY id DESC LIMIT 5;`)
	if err != nil {
		log.Printf("Unable to get data (Order): %vn", err)
	}
	for rows.Next() {
		var oid string
		err = rows.Scan(&oid)
		if err != nil {
			log.Printf("Unable to get data (Order): %vn", err)
		}
		order, _ := getdb(oid)
		cache.Store(oid, order)
	}
}

func getdb(oid string) (Order, int) {
	config := map[string]string{
		"host":     "localhost",
		"port":     "5432",
		"user":     "postgres",
		"password": "1",
		"database": "L",
	}
	pool, err := ConnectToPostgres(config)
	if err != nil {
		fmt.Println("не успешко")
	}
	defer pool.Close()
	var o Order
	err = pool.QueryRow(context.Background(), `SELECT order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, delivery_service_1, sm_id, date_created, oof_shard
	FROM public."order" where order_uid = $1`, oid).Scan(&o.OrderUID, &o.TrackNumber, &o.Entry, &o.Locale, &o.InternalSignature, &o.CustomerID, &o.Delivery_Service, &o.Shardkey, &o.SmID, &o.Date_Created, &o.Oof_shard)
	if err != nil {
		log.Printf("Unable to get Order from database: %v\n", err)
		return o, -1
	}
	err = pool.QueryRow(context.Background(), `SELECT request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee, order_order_uid, order_track_number
	FROM public.payment where order_order_uid = $1`, oid).Scan(&o.Payment.RequestID, &o.Payment.Currency, &o.Payment.Provider, &o.Payment.Amount, &o.Payment.PaymentDt, &o.Payment.Bank, &o.Payment.DeliveryCost, &o.Payment.GoodsTotal, &o.Payment.CustomFee, &o.Payment.OrderUID, &o.Payment.TrackNumber)
	if err != nil {
		log.Printf("Unable to get payment from database: %v\n", err)
		return o, -2
	}
	rowsItems, err := pool.Query(context.Background(), `SELECT order_track_number, chrt_id, price, rid, name, sale, size, total_price, nm_id, brand, status, order_order_uid, order_track_number
	FROM "order"
	INNER JOIN items ON items.order_track_number = "order".track_number
	where order_order_uid = $1`, oid)
	if err != nil {
		log.Printf("Unable to get items from database: %v\n", err)
		return o, -3
	}
	defer rowsItems.Close()
	rowc := 0
	err = pool.QueryRow(context.Background(), `select count(*) from (SELECT order_track_number, chrt_id, price, rid, name, sale, size, total_price, nm_id, brand, status, order_order_uid, order_track_number
		FROM "order"
		INNER JOIN items ON items.order_track_number = "order".track_number
		where order_order_uid = $1)`, oid).Scan(&rowc)
	if err != nil {
		log.Printf("Unable to get items from database: %v\n", err)
		return o, -4
	}
	o.Items = make([]Items, rowc)
	i := 0
	for rowsItems.Next() {
		err := rowsItems.Scan(&o.TrackNumber, &o.Items[i].ChrtID, &o.Items[i].Price, &o.Items[i].Rid, &o.Items[i].Name, &o.Items[i].Sale, &o.Items[i].Size, &o.Items[i].TotalPrice, &o.Items[i].NmID, &o.Items[i].Brand, &o.Items[i].Status, &o.Items[i].OrderUID, &o.Items[i].TrackNumber)
		if err != nil {
			log.Printf("Unable to scan row: %vn", err)
			continue
		}
		i++
	}
	err = pool.QueryRow(context.Background(), `SELECT name, phone, zip, city, address, region, email, order_order_uid, order_track_number
	FROM public.delivery where order_order_uid = $1`, oid).Scan(&o.Delivery.Name, &o.Delivery.Phone, &o.Delivery.Zip, &o.Delivery.City, &o.Delivery.Address, &o.Delivery.Region, &o.Delivery.Email, &o.Delivery.OrderUID, &o.Delivery.TrackNumber)
	if err != nil {
		log.Printf("Unable to get delivery from database: %v\n", err)
		return o, -5
	}
	fmt.Println("Мною воспользовались", oid)
	return o, 0
}

func insertdb(o Order) {
	config := map[string]string{
		"host":     "localhost",
		"port":     "5432",
		"user":     "postgres",
		"password": "1",
		"database": "L",
	}
	pool, err := ConnectToPostgres(config)
	if err != nil {
		fmt.Println("не успешко")
	} else {
		fmt.Println("успешко")
	}
	defer pool.Close()
	var lastInsertId int
	err = pool.QueryRow(context.Background(), `INSERT INTO public."order"(
		order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, delivery_service_1, sm_id, date_created, oof_shard)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);`, o.OrderUID, o.TrackNumber, o.Entry, o.Locale, o.InternalSignature, o.CustomerID, o.Delivery_Service, o.Shardkey, o.SmID, o.Date_Created, o.Oof_shard).Scan(&lastInsertId)
	if err != nil {
		errStr := err.Error()
		if errStr == "no rows in result set" {
			fmt.Println("good")
		} else {
			log.Printf("Unable to insert data (Order): %v\n", err)
		}
	}
	for _, item := range o.Items {
		err := pool.QueryRow(context.Background(), `INSERT INTO public.items(
			chrt_id, price, rid, name, sale, size, total_price, nm_id, brand, status, order_order_uid, order_track_number)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);`, item.ChrtID, item.Price, item.Rid, item.Name, item.Sale, item.Size,
			item.TotalPrice, item.NmID, item.Brand, item.Status, o.OrderUID, o.TrackNumber).Scan(&lastInsertId)
		if err != nil {
			errStr := err.Error()
			if errStr == "no rows in result set" {
				fmt.Println("Item - good")
			} else {
				log.Printf("Unable to insert data (Item): %v\n", err)
			}
		}
	}
	err = pool.QueryRow(context.Background(), `INSERT INTO public.payment(
		request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee, order_order_uid, order_track_number)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);`, o.Payment.RequestID, o.Payment.Currency, o.Payment.Provider, o.Payment.Amount, o.Payment.PaymentDt, o.Payment.Bank, o.Payment.DeliveryCost, o.Payment.GoodsTotal, o.Payment.CustomFee, o.OrderUID, o.TrackNumber).Scan(&lastInsertId)
	if err != nil {
		errStr := err.Error()
		if errStr == "no rows in result set" {
			fmt.Println("Payment - good")
		} else {
			log.Printf("Unable to insert data (Payment): %v\n", err)
		}
	}
	err = pool.QueryRow(context.Background(), `INSERT INTO public.delivery(
		name, phone, zip, city, address, region, email, order_order_uid, order_track_number)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);`, o.Delivery.Name, o.Delivery.Phone, o.Delivery.Zip, o.Delivery.City, o.Delivery.Address, o.Delivery.Region, o.Delivery.Email, o.OrderUID, o.TrackNumber).Scan(&lastInsertId)
	if err != nil {
		errStr := err.Error()
		if errStr == "no rows in result set" {
			fmt.Println("Delivery - good")
		} else {
			log.Printf("Unable to insert data (Delivery): %v\n", err)
		}
	}
}

func main() {
	sc, err := stan.Connect("test-cluster", "client-1", stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sc.Close()
	chanMsg := make(chan *stan.Msg, 10)
	sub, err := sc.Subscribe("rabotai", func(msg *stan.Msg) {
		chanMsg <- msg
	}, stan.DurableName("my_durable"), stan.MaxInflight(5))
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Close()
	go func() {
		for msg := range chanMsg {
			var order Order
			if err := json.Unmarshal(msg.Data, &order); err != nil {
				log.Printf("json.Unmarshal error: %v\n", err)
			}
			insertdb(order)
			msg.Ack()
		}
	}()
	restoreCache()
	http.HandleFunc("/order/", order)
	http.ListenAndServe(":8080", nil)
}

func order(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/order/"):]
	log.Println(r.URL.Path[len("/order"):])
	order, i := get(id)
	if i != 0 {
		http.Error(w, "Такого заказа не существует", http.StatusNotFound)
		return
	}
	json, err := json.Marshal(order)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
