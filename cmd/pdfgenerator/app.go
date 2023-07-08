package main

import (
	"github.com/jung-kurt/gofpdf"
	"log"
)

var _ PdfDocument = &pdfDocument{}

type pdf struct {
	*gofpdf.Fpdf
}

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
	OutputFileAndClose()
	Header(string)
	AddText(string)
	AddTextRight(string)
	AddCheckBox(float64, string)
	TableHeader(string, string)
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
	p.pdf.CellFormat(190, p.lineHeight, text, "0", 0, "C", false, 0, "") //вывод текста
}

// Верхний блок
func (p *pdfDocument) AddText(text string) {
	p.pdf.SetFont("Arial", "", 14) //шрифт,жирность,размер
	_, lineHt := p.pdf.GetFontSize()
	p.pdf.Ln(lineHt * 2)

	p.pdf.CellFormat(95, lineHt*3, text, "0", 0, "L", false, 0, "")
}

// Верхний блок,правая строка
func (p *pdfDocument) AddTextRight(text string) {
	p.pdf.SetFont("Arial", "B", 14)
	_, lineHt := p.pdf.GetFontSize()
	p.pdf.SetTextColor(52, 116, 178) //цвет текста

	p.pdf.CellFormat(95, lineHt*3, text, "0", 0, "R", false, 0, "")
	p.pdf.SetTextColor(0, 0, 0)
}

// AddCheckBox
func (p *pdfDocument) AddCheckBox(width float64, text string) {
	p.pdf.SetFont("Arial", "", 10)
	_, lineHt := p.pdf.GetFontSize()
	p.pdf.CellFormat(width, lineHt*2, text, "1", 0, "C", false, 0, "")
}

// заголовок таблиц
func (p *pdfDocument) TableHeader(text, alignStr string) {
	p.pdf.SetFillColor(52, 116, 178)  // Установка цвета заливки для заголовка
	p.pdf.SetTextColor(255, 255, 255) // Устанавливает цвет текста
	p.pdf.SetDrawColor(227, 227, 227) // Устанавливаем цвет границы в синий

	widthTable := 48.5
	heightTable := 14.0
	x, y := p.pdf.GetXY() // получение текущих координат X и Y

	if text == " Jednotková \n cena bez DPH " {
		p.pdf.MultiCell(widthTable, 7, text, "1", alignStr, true)
	} else {
		p.pdf.MultiCell(widthTable, heightTable, text, "1", alignStr, true)
	}
	p.pdf.SetXY(x+widthTable, y) // установка новых координат X и Y, увеличиваем X
}

// Создание pdf
func (p *pdfDocument) OutputFileAndClose() {
	err := p.pdf.OutputFileAndClose(p.outputFileAndClose)
	if err != nil {
		log.Println(err)
		return
	}
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

	pdf := newPdfDocument()
	pdf.AddPage()
	pdf.AddUTF8Font() // Добавляем шрифт, поддерживающий больше символов
	pdf.SetFont()     // Задаем шрифт, жирность и размер

	//*Заголовок документа
	pdf.Header("F A K T U R A")
	pdf.LineHt(1)

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

	//***
	//*Создаем pdf файл
	pdf.OutputFileAndClose()

	////*Таблица
	////Header
	//pdf.Ln(10)
	//pdf.SetFillColor(52, 116, 178)  // Установка цвета заливки для заголовка
	//pdf.SetTextColor(255, 255, 255) // Устанавливает цвет текста
	//pdf.SetDrawColor(227, 227, 227) // Устанавливаем цвет границы в синий
	//
	//widthTable := 48.5
	//heightTable := 14.0
	//pdf.CellFormat(widthTable, heightTable, " Popis/Výkon ", "1", 0, "L", true, 0, "")
	//pdf.CellFormat(widthTable, heightTable, "Množství", "1", 0, "C", true, 0, "")
	//pdf.CellFormat(widthTable, heightTable, "Cena za kus", "1", 0, "C", true, 0, "")
	//pdf.MultiCell(widthTable, 7, " Jednotková \n cena bez DPH ", "1", "C", true)
	//
	////Body
	//pdf.Ln(0)
	//pdf.SetFillColor(255, 255, 255) // Установка цвета заливки для заголовка
	//pdf.SetTextColor(0, 0, 0)       // Устанавливает цвет текста
	//
	//pdf.CellFormat(widthTable, heightTable, " Vyjezd ", "1", 0, "L", true, 0, "")
	//pdf.CellFormat(widthTable, heightTable, "2", "1", 0, "C", true, 0, "")
	//pdf.CellFormat(widthTable, heightTable, "600.00", "1", 0, "C", true, 0, "")
	//pdf.CellFormat(widthTable, heightTable, "1200.00", "1", 0, "C", true, 0, "")
	//
	//pdf.Ln(14)
	//pdf.CellFormat(widthTable, heightTable, " Odmitani ", "1", 0, "L", true, 0, "")
	//pdf.CellFormat(widthTable, heightTable, "1", "1", 0, "C", true, 0, "")
	//pdf.CellFormat(widthTable, heightTable, "2600.00", "1", 0, "C", true, 0, "")
	//pdf.CellFormat(widthTable, heightTable, "2600.00", "1", 0, "C", true, 0, "")
	//
	////*Нижний блок
	////line 1
	//pdf.Ln(20)
	//widthDownBlock := 48.5
	//pdf.CellFormat(widthDownBlock, lineHt+2, " Částka obdržena: ", "0", 0, "L", false, 0, "")
	//pdf.CellFormat(widthDownBlock, lineHt+2, "  ", "0", 0, "C", false, 0, "")
	//pdf.CellFormat(widthDownBlock+20.5, lineHt+2, " Mezisoučet: ", "0", 0, "R", false, 0, "")
	//pdf.CellFormat(28, lineHt+2, " 3800.00 ", "0", 0, "R", false, 0, "")
	////line 2
	//pdf.Ln(lineHt + 2)
	//pdf.CellFormat(widthDownBlock, lineHt+2, " Hotově: ", "0", 0, "L", false, 0, "")
	//pdf.CellFormat(widthDownBlock, lineHt+2, " 0.00 ", "0", 0, "C", false, 0, "")
	//pdf.CellFormat(widthDownBlock+20.5, lineHt+2, "  ", "0", 0, "R", false, 0, "")
	//pdf.CellFormat(28, lineHt+2, "  ", "0", 0, "R", false, 0, "")
	////line 3
	//pdf.Ln(lineHt + 2)
	//pdf.CellFormat(widthDownBlock, lineHt+2, " Kartou: ", "0", 0, "L", false, 0, "")
	//pdf.CellFormat(widthDownBlock, lineHt+2, " 0.00 ", "0", 0, "C", false, 0, "")
	//pdf.CellFormat(widthDownBlock+20.5, lineHt+2, " DPH 21 %: ", "0", 0, "R", false, 0, "")
	//pdf.CellFormat(28, lineHt+2, " 798.00 ", "0", 0, "R", false, 0, "")
	////line 4
	//pdf.Ln(lineHt + 2)
	//pdf.CellFormat(widthDownBlock, lineHt+2, " Převod: ", "0", 0, "L", false, 0, "")
	//pdf.CellFormat(widthDownBlock, lineHt+2, " 0.00 ", "0", 0, "C", false, 0, "")
	//pdf.CellFormat(widthDownBlock+20.5, lineHt+2, "  ", "0", 0, "R", false, 0, "")
	//pdf.CellFormat(28, lineHt+2, "  ", "0", 0, "R", false, 0, "")
	////line 5
	//pdf.Ln(lineHt + 2)
	//pdf.CellFormat(widthDownBlock, lineHt+2, " Dluh: ", "0", 0, "L", false, 0, "")
	//pdf.CellFormat(widthDownBlock, lineHt+2, " 4598.00 ", "0", 0, "C", false, 0, "")
	//pdf.CellFormat(widthDownBlock+20.5, lineHt+2, " Celková částka: ", "0", 0, "R", false, 0, "")
	//pdf.CellFormat(28, lineHt+2, " 4598.00 ", "0", 0, "R", false, 0, "")
	//
	////*Нижний блок
	//// Задаем размер нижнего поля
	//bottomMargin := 26.0
	//// Получаем размеры страницы
	//_, pageHeight := pdf.GetPageSize()
	//// Устанавливаем новую высоту Y, вычитая нижний отступ и высоту строки из высоты страницы
	//pdf.SetY(pageHeight - bottomMargin - lineHt)
	//pdf.SetFont("Arial", "", 11) // Установка шрифта перед выводом текста
	//pdf.MultiCell(190, lineHt, "Rychly servis bohemia 24/7 s.r.o, IČO 17973538, Braunerova 563/7, Libeň, 180 00 Praha 8\nBankovní účet: 5040636073/0800", "0", "C", false)

	//*Создаем pdf файл
	//err := pdf.OutputFileAndClose("../../yourContract.pdf")
	//if err != nil {
	//	panic(err)
	//}
}
