package main

import (
	"fmt"
)

type Resource interface { // как элемент проекта
	Connect() error
	Disconnect() error
	SendData(data string) error
	ReceiveData() (string, error)
}

type DatabaseAdapter struct {
	dbConfig string
}

func (db *DatabaseAdapter) Connect() error {
	fmt.Println("Подключение к БД")
	return nil
}

func (db *DatabaseAdapter) Disconnect() error {
	fmt.Println("Отключение от БД")
	return nil
}

func (db *DatabaseAdapter) SendData(data string) error {
	fmt.Printf("Отправляем данные '%s' в БД\n", data)
	return nil
}

func (db *DatabaseAdapter) ReceiveData() (string, error) {
	fmt.Println("Данные из БД")
	return "Данные из БД", nil
}

type BotAdapter struct {
	botConfig string
}

func (bot *BotAdapter) Connect() error {
	fmt.Println("Подключение к боту")
	return nil
}

func (bot *BotAdapter) Disconnect() error {
	fmt.Println("Отключение от бота")
	return nil
}

func (bot *BotAdapter) SendData(data string) error {
	fmt.Printf("Отправляем данные '%s' боту\n", data)
	return nil
}

func (bot *BotAdapter) ReceiveData() (string, error) {
	fmt.Println("Данные от бота")
	return "Сообщение от бота", nil
}

type WebAdapter struct {
	websiteConfig string
}

func (site *WebAdapter) Connect() error {
	fmt.Println("Подключение к сайту")
	return nil
}

func (site *WebAdapter) Disconnect() error {
	fmt.Println("Отключение от сайта")
	return nil
}

func (site *WebAdapter) SendData(data string) error {
	fmt.Printf("Отправляем данные '%s' на сайт\n", data)
	return nil
}

func (site *WebAdapter) ReceiveData() (string, error) {
	fmt.Println("Данные с сайта")
	return "Данные с сайта", nil
}

func main() {
	projects := [3]Resource{&DatabaseAdapter{"db_config"}, &BotAdapter{"bot_config"}, &WebAdapter{"website_config"}} // инициализируем 3 проекта

	for _, project := range projects { // проходимся по каждому и пробуем все команды для проверки: "живо все или нет?"
		err := project.Connect()
		if err != nil {
			fmt.Println("Ошибка подключения:", err)
			continue
		}

		err = project.SendData("Отправка тестового сообщения!")
		if err != nil {
			fmt.Println("Ошибка отправки данных:", err)
			continue
		}

		data, err := project.ReceiveData()
		if err != nil {
			fmt.Println("Ошибка получения данных:", err)
			continue
		}
		fmt.Println("Полученные данные:", data)

		err = project.Disconnect()
		if err != nil {
			fmt.Println("Ошибка отключения:", err)
			continue
		}
	}
}
