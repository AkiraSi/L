package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/nats-io/stan.go"
)

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

func main() {
	sc, err := stan.Connect("test-cluster", "client-2", stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sc.Close()
	item1 := Items{ChrtID: 1, Price: 10, Rid: "rid 1", Name: "T-Shirt-4", Sale: 9, Size: "M", TotalPrice: 13, NmID: 1, Brand: "Adidas", Status: 123, OrderUID: "o2", TrackNumber: "t2"}
	item2 := Items{ChrtID: 2, Price: 12, Rid: "rid 2", Name: "Jeans", Sale: 11, Size: "S", TotalPrice: 14, NmID: 2, Brand: "Collins", Status: 23, OrderUID: "o2", TrackNumber: "t2"}
	item3 := Items{ChrtID: 3, Price: 18, Rid: "rid 3", Name: "Sneakers", Sale: 15, Size: "M", TotalPrice: 20, NmID: 1, Brand: "Nike", Status: 12133, OrderUID: "o2", TrackNumber: "t2"}
	payment := Payment{Currency: "Rub", Provider: "Provider 1", Amount: 47, PaymentDt: 2, Bank: "VTB", DeliveryCost: 7, GoodsTotal: 3, OrderUID: "o2", TrackNumber: "t2", RequestID: "ri", CustomFee: 0}
	delivery := Delivery{Name: "Name", Phone: "Phone", Zip: "Zip", City: "City", Address: "Address", Region: "Region", Email: "Email", OrderUID: "o2", TrackNumber: "t2"}
	order := Order{OrderUID: "o35", Entry: "2", InternalSignature: "IS 2", Payment: payment, Items: []Items{item1, item2, item3},
		Locale: "Ru", CustomerID: "2", TrackNumber: "t35", Delivery_Service: "DS 2", Shardkey: "SK 2", SmID: 2, Date_Created: "123", Oof_shard: "shard", Delivery: delivery}
	orderData, err := json.Marshal(order)
	if err != nil {
		log.Printf("json.Marshal error: %v\n", err)
	}
	err = sc.Publish("rabotai", orderData)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Message published")
}
