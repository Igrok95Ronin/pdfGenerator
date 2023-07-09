package main

import (
	"encoding/json"
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"io"
	"log"
	"net/http"
	"time"
)

var _ PdfDocument = &pdfDocument{}

// Структура Json
type dataJsonStruct struct {
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
		Name     string `json:"Name"`
		Amount   int    `json:"Amount"`
		Price    int    `json:"Price"`
		PriceBuy int    `json:"PriceBuy"`
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
func (d *dataJsonStruct) Parse(url string) {
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

// Структура PDF
type pdf struct {
	*gofpdf.Fpdf
}

// Структура PDF
type pdfDocument struct {
	pdf
	fontFamily         string
	fontStyle          string
	fontSize           float64
	fontFilePath       string
	lineHeight         float64
	outputFileAndClose string
}

type PdfDocument interface {
	AddPage()
	SetFont()
	AddUTF8Font()
	LineHt(float64)
	OutputFileAndClose() error
	Header(string)
	AddText(string)
	AddTextRight(string)
	AddCheckBox(float64, string)
	TableHeader(string, string)
	TableBody(string, string)
	BottomBlock(float64, string, string)
	Footer(string)
}

func newPdfDocument() PdfDocument {
	return &pdfDocument{
		pdf{
			gofpdf.New("P", "mm", "A4", ""),
		},
		"Arial",
		"B",
		16,
		"../../ui/static/fonts/DejaVuSans.ttf",
		16,
		"../../yourContract.pdf",
	}
}

// Заголовок документа
func (p *pdfDocument) Header(text string) {
	p.pdf.AddUTF8Font("DejaVuSans", "B", "../../ui/static/fonts/DejaVuSans-Bold.ttf")
	p.pdf.SetFont("DejaVuSans", "B", 16)                                 //шрифт,жирность,размер
	p.pdf.CellFormat(190, p.lineHeight, text, "0", 0, "C", false, 0, "") //вывод текста
}

// Верхний блок
func (p *pdfDocument) AddText(text string) {
	p.pdf.SetFont("Arial", "", 12) //шрифт,жирность,размер
	_, lineHt := p.pdf.GetFontSize()
	p.pdf.Ln(lineHt * 1.5)

	p.pdf.CellFormat(95, lineHt*3, text, "0", 0, "L", false, 0, "")
}

// Верхний блок,правая строка
func (p *pdfDocument) AddTextRight(text string) {
	p.pdf.AddUTF8Font("DejaVuSans", "B", "../../ui/static/fonts/DejaVuSans-Bold.ttf")
	p.pdf.SetFont("DejaVuSans", "B", 12)
	_, lineHt := p.pdf.GetFontSize()
	p.pdf.SetTextColor(52, 116, 178) //цвет текста

	p.pdf.CellFormat(95, lineHt*3, text, "0", 0, "R", false, 0, "")
	p.pdf.SetTextColor(0, 0, 0)
}

// AddCheckBox
func (p *pdfDocument) AddCheckBox(width float64, text string) {
	p.pdf.SetFont("Arial", "", 10)
	_, lineHt := p.pdf.GetFontSize()
	p.pdf.CellFormat(width, lineHt*2, text, "0", 0, "C", false, 0, "")
}

// заголовок таблиц
func (p *pdfDocument) TableHeader(text, alignStr string) {
	p.pdf.SetFillColor(52, 116, 178)  // Установка цвета заливки для заголовка
	p.pdf.SetTextColor(255, 255, 255) // Устанавливает цвет текста
	p.pdf.SetDrawColor(227, 227, 227) // Устанавливаем цвет границы в синий

	p.pdf.AddUTF8Font("DejaVuSans", "B", "../../ui/static/fonts/DejaVuSans-Bold.ttf")
	p.pdf.SetFont("DejaVuSans", "B", 10) //шрифт,жирность,размер

	const (
		WidthTable  = 48.5
		HeightTable = 14
	)
	x, y := p.pdf.GetXY() // получение текущих координат X и Y

	if text == " Jednotková \n cena bez DPH " {
		p.pdf.MultiCell(WidthTable, 7, text, "1", alignStr, true)
	} else {
		p.pdf.MultiCell(WidthTable, HeightTable, text, "1", alignStr, true)
	}
	p.pdf.SetXY(x+WidthTable, y) // установка новых координат X и Y, увеличиваем X
}

// тело таблицы
func (p *pdfDocument) TableBody(text, alignStr string) {
	widthTable := 48.5
	heightTable := 14.0
	p.pdf.SetFillColor(255, 255, 255) // Установка цвета заливки для заголовка
	p.pdf.SetTextColor(0, 0, 0)       // Устанавливает цвет текста
	p.pdf.CellFormat(widthTable, heightTable, text, "1", 0, alignStr, true, 0, "")
}

// Нижний блок
func (p *pdfDocument) BottomBlock(width float64, text, alignStr string) {
	//widthDownBlock := 48.5
	_, lineHt := p.pdf.GetFontSize()
	p.pdf.CellFormat(width, lineHt+2, text, "0", 0, alignStr, false, 0, "")
}

// Footer
func (p *pdfDocument) Footer(text string) {
	// Задаем размер нижнего поля
	bottomMargin := 28.0
	// Получаем размеры страницы
	_, pageHeight := p.pdf.GetPageSize()
	p.pdf.SetFont("Arial", "", 10) // Установка шрифта перед выводом текста
	_, lineHt := p.pdf.GetFontSize()
	// Устанавливаем новую высоту Y, вычитая нижний отступ и высоту строки из высоты страницы
	p.pdf.SetY(pageHeight - bottomMargin - lineHt)
	p.pdf.MultiCell(190, lineHt*1.5, text, "0", "C", false)

}

// Создание pdf
func (p *pdfDocument) OutputFileAndClose() error {
	return p.pdf.OutputFileAndClose(p.outputFileAndClose)
}

func (p *pdfDocument) AddPage() {
	p.pdf.AddPage()
}

func (p *pdfDocument) SetFont() {
	p.pdf.SetFont(p.fontFamily, p.fontStyle, p.fontSize)
}

func (p *pdfDocument) AddUTF8Font() {
	p.pdf.AddUTF8Font(p.fontFamily, "", p.fontFilePath)
}

// Перенос на новую строчку
func (p *pdfDocument) LineHt(ht float64) {
	p.pdf.SetFont("Arial", "", 10)
	_, lineHt := p.pdf.GetFontSize()
	p.pdf.Ln(lineHt * ht) //перенос строки
}

// test
func main() {
	jsn := dataJsonStruct{}
	jsn.Parse("https://app.o95.info/receipt?secret=secret72353&id=205887")
	fmt.Println(jsn.Amount)

	//PDF
	pdf := newPdfDocument()
	pdf.AddPage()
	pdf.AddUTF8Font() // Добавляем шрифт, поддерживающий больше символов
	pdf.SetFont()     // Задаем шрифт, жирность и размер

	//*Заголовок документа
	pdf.Header("F A K T U R A")
	pdf.LineHt(2)

	//*Верхний блок
	//-Первая строка
	pdf.AddText("Kamil Teplý")

	//-Правый строка ID
	pdf.AddTextRight("ID: CZ-3135")

	//-Вторая строка
	pdf.AddText("Francouzská 2")

	//-Третья строка
	pdf.AddText("12000 Praha")
	pdf.LineHt(7)

	//*CheckBox
	pdf.AddCheckBox(38, "[✓] Objednavka")
	pdf.AddCheckBox(25, "[ ] Nabidka")
	pdf.AddCheckBox(37, "[ ] Konzultace")
	pdf.AddCheckBox(37, "[✓] Nalehavost")
	pdf.AddCheckBox(25, "[ ] Montaz")
	pdf.AddCheckBox(31.6, "[✓] Pojisteni")
	pdf.LineHt(3)

	//*Таблица
	//-Header
	pdf.TableHeader(" Popis/Výkon ", "L")
	pdf.TableHeader(" Množství ", "C")
	pdf.TableHeader(" Cena za kus ", "C")
	pdf.TableHeader(" Jednotková \n cena bez DPH ", "C")
	pdf.LineHt(4)

	//-Body
	pdf.TableBody(" Vyjezd ", "L")
	pdf.TableBody(" 2 ", "C")
	pdf.TableBody(" 600.00 ", "C")
	pdf.TableBody(" 1200.00 ", "C")
	pdf.LineHt(4)
	pdf.TableBody(" Odmitani ", "L")
	pdf.TableBody(" 1 ", "C")
	pdf.TableBody(" 2600.00 ", "C")
	pdf.TableBody(" 2600.00 ", "C")
	pdf.LineHt(6)

	//*Нижний блок
	//-line 1
	pdf.BottomBlock(48.5, "Částka obdržena:", "L")
	pdf.BottomBlock(48.5, "", "L")
	pdf.BottomBlock(69, "Mezisoučet:", "R")
	pdf.BottomBlock(28, "3800.00", "R")
	pdf.LineHt(2)

	//-line 2
	pdf.BottomBlock(48.5, "Hotově:", "L")
	pdf.BottomBlock(48.5, "0.00", "L")
	pdf.BottomBlock(69, "", "R")
	pdf.BottomBlock(28, "", "R")
	pdf.LineHt(2)

	//-line 3
	pdf.BottomBlock(48.5, "Kartou:", "L")
	pdf.BottomBlock(48.5, "0.00", "L")
	pdf.BottomBlock(69, "DPH 21 %:", "R")
	pdf.BottomBlock(28, "798.00", "R")
	pdf.LineHt(2)

	//-line 4
	pdf.BottomBlock(48.5, "Převod:", "L")
	pdf.BottomBlock(48.5, "0.00", "L")
	pdf.BottomBlock(69, "", "R")
	pdf.BottomBlock(28, "", "R")
	pdf.LineHt(2)

	//-line 5
	pdf.BottomBlock(48.5, "Dluh:", "L")
	pdf.BottomBlock(48.5, "4598.00", "L")
	pdf.BottomBlock(69, "Celková částka:", "R")
	pdf.BottomBlock(28, "4598.00", "R")

	//*Footer
	pdf.Footer("Rychly servis bohemia 24/7 s.r.o, IČO 17973538, Braunerova 563/7, Libeň, 180 00 Praha 8\nBankovní účet: 5040636073/0800")

	//*Создаем pdf файл
	err := pdf.OutputFileAndClose()
	if err != nil {
		log.Fatalf("Не удалось вывести файл и закрыть PDF-документ: %v", err)
	}

}
