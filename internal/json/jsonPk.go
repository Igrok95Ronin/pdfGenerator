package jsonPk

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

// Структура Json
type DataJsonStruct struct {
	Amount       string    `json:"Amount"`
	AmountATM    string    `json:"AmountATM"`
	AmountCash   string    `json:"AmountCash"`
	AmountCredit string    `json:"AmountCredit"`
	AmountOnline string    `json:"AmountOnline"`
	AmountTax    string    `json:"AmountTax"`
	AmountTotal  string    `json:"AmountTotal"`
	CheckBox1    string    `json:"CheckBox1"`
	CheckBox2    string    `json:"CheckBox2"`
	CheckBox3    string    `json:"CheckBox3"`
	CheckBox4    string    `json:"CheckBox4"`
	CheckBox5    string    `json:"CheckBox5"`
	CheckBox6    string    `json:"CheckBox6"`
	City         string    `json:"City"`
	CompletedAt  time.Time `json:"CompletedAt"`
	Country      string    `json:"Country"`
	CountryTax   float64   `json:"CountryTax"`
	CreatedAt    time.Time `json:"CreatedAt"`
	EmployeeFirm string    `json:"EmployeeFirm"`
	Expenses     []struct {
		Name     string  `json:"Name"`
		Amount   int     `json:"Amount"`
		Price    float64 `json:"Price"`
		PriceBuy float64 `json:"PriceBuy"`
	} `json:"Expenses"`
	FakeCompany     bool      `json:"FakeCompany"`
	Firma           string    `json:"Firma"`
	ID              int       `json:"ID"`
	Name            string    `json:"Name"`
	Note            string    `json:"Note"`
	PassportNumber  string    `json:"PassportNumber"`
	Phone           string    `json:"Phone"`
	PropertyID      string    `json:"PropertyID"`
	RID             string    `json:"RID"`
	Radio1          string    `json:"Radio1"`
	Radio2          string    `json:"Radio2"`
	Radio3          string    `json:"Radio3"`
	Radio4          string    `json:"Radio4"`
	Radio5          string    `json:"Radio5"`
	SignatureEndURL string    `json:"SignatureEndURL"`
	SignatureURL    string    `json:"SignatureURL"`
	StartedAt       time.Time `json:"StartedAt"`
	Street          string    `json:"Street"`
	Type            string    `json:"Type"`
	ZipCode         string    `json:"ZipCode"`
}

// Декодирование JSON
func (d *DataJsonStruct) Parse(url string) {
	// URL, откуда мы хотим получить JSON

	// Выполнение HTTP GET запроса
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Ошибка отправки запроса: %s", err)
	}
	defer resp.Body.Close() // Не забываем закрыть тело ответа после чтения

	// Чтение всего тела ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Ошибка чтения тела ответа: %s", err)
	}

	// Декодирование JSON из тела ответа в нашу структуру
	err = json.Unmarshal(body, d)
	if err != nil {
		log.Fatalf("Ошибка при разборе JSON: %s", err)
	}

	// Вывод декодированных данных
	//fmt.Printf("Decoded data: %+v\n", d)
}
