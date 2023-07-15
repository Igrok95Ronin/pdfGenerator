package pdf

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"log"
	"pdgGenerator/internal/json"
	"strconv"
	"unicode/utf8"
)

var _ PdfDocument = &pdfDocument{}

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
	TableHeader(float64, float64, string, string)
	TableBody(float64, float64, string, string)
	BottomBlock(float64, string, string)
	Footer(string)
	SecondLeaf(string)
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
func (p *pdfDocument) TableHeader(width, height float64, text, alignStr string) {
	p.pdf.SetFillColor(52, 116, 178)  // Установка цвета заливки для заголовка
	p.pdf.SetTextColor(255, 255, 255) // Устанавливает цвет текста
	p.pdf.SetDrawColor(227, 227, 227) // Устанавливаем цвет границы в синий

	p.pdf.AddUTF8Font("DejaVuSans", "B", "../../ui/static/fonts/DejaVuSans-Bold.ttf")
	p.pdf.SetFont("DejaVuSans", "B", 10) //шрифт,жирность,размер

	x, y := p.pdf.GetXY() // получение текущих координат X и Y

	if text == " Jednotková \n cena bez DPH " {
		p.pdf.MultiCell(width, 7, text, "1", alignStr, true)
	} else {
		p.pdf.MultiCell(width, height, text, "1", alignStr, true)
	}
	p.pdf.SetXY(x+width, y) // установка новых координат X и Y, увеличиваем X
}

// тело таблицы
func (p *pdfDocument) TableBody(width, height float64, text, alignStr string) {
	x, y := p.pdf.GetXY() // получение текущих координат X и Y

	p.pdf.SetFillColor(255, 255, 255) // Установка цвета заливки для заголовка
	p.pdf.SetTextColor(0, 0, 0)       // Устанавливает цвет текста
	p.pdf.MultiCell(width, height, text, "1", alignStr, true)

	p.pdf.SetXY(x+width, y) // установка новых координат X и Y, увеличиваем X
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

// Второй лист
func (p *pdfDocument) SecondLeaf(text string) {
	_, lineHt := p.pdf.GetFontSize()
	p.pdf.CellFormat(190, lineHt+20, text, "0", 0, "L", false, 0, "")
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

func Pdf(url string) {

	jsn := &json.DataJsonStruct{}
	jsn.Parse(url)
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
	pdf.AddText(jsn.Name)

	//-Правый строка ID
	pdf.AddTextRight("ID: " + jsn.RID)

	//-Вторая строка
	pdf.AddText(jsn.Street + " " + jsn.PropertyID)

	//-Третья строка
	pdf.AddText(jsn.ZipCode + " " + jsn.City)
	pdf.LineHt(7)

	//*CheckBox

	if jsn.CheckBox1 == "yes" {
		pdf.AddCheckBox(37.5, "[✓] Objednavka")
	} else {
		pdf.AddCheckBox(37.5, "[ ] Objednavka")
	}
	if jsn.CheckBox2 == "yes" {
		pdf.AddCheckBox(24.5, "[✓] Nabidka")
	} else {
		pdf.AddCheckBox(24.5, "[ ] Nabidka")
	}
	if jsn.CheckBox3 == "yes" {
		pdf.AddCheckBox(36.5, "[✓] Konzultace")
	} else {
		pdf.AddCheckBox(36.5, "[ ] Konzultace")
	}
	if jsn.CheckBox4 == "yes" {
		pdf.AddCheckBox(36.5, "[✓] Nalehavost")
	} else {
		pdf.AddCheckBox(36.5, "[ ] Nalehavost")
	}
	if jsn.CheckBox5 == "yes" {
		pdf.AddCheckBox(24.5, "[✓] Montaz")
	} else {
		pdf.AddCheckBox(24.5, "[ ] Montaz")
	}
	if jsn.CheckBox6 == "yes" {
		pdf.AddCheckBox(30.5, "[✓] Pojisteni")
	} else {
		pdf.AddCheckBox(30.5, "[ ] Pojisteni")
	}
	pdf.LineHt(3)

	//*Таблица
	//-Header
	pdf.TableHeader(55.0, 14, " Popis/Výkon ", "L")
	pdf.TableHeader(45.0, 14, " Množství ", "C")
	pdf.TableHeader(45.0, 14, " Cena za kus ", "C")
	pdf.TableHeader(45.0, 14, " Jednotková \n cena bez DPH ", "C")
	pdf.LineHt(4)

	//-Body
	for i := 0; i < len(jsn.Expenses); i++ {
		if utf8.RuneCountInString(jsn.Expenses[i].Name) > 30 {
			pdf.TableBody(55.0, 7, jsn.Expenses[i].Name, "L")
		} else {
			pdf.TableBody(55.0, 14, jsn.Expenses[i].Name, "L")
		}
		pdf.TableBody(45.0, 14, strconv.FormatFloat(jsn.Expenses[i].Amount, 'f', -1, 64), "C")
		pdf.TableBody(45.0, 14, strconv.FormatFloat(jsn.Expenses[i].Price, 'f', -1, 64), "C")
		pdf.TableBody(45.0, 14, strconv.FormatFloat(jsn.Expenses[i].PriceBuy, 'f', -1, 64), "C")
		pdf.LineHt(4)
	}
	pdf.LineHt(2)

	//*Нижний блок
	//-line 1
	const (
		width1 = 47.5
		width2 = 68
		width3 = 27
	)
	pdf.BottomBlock(width1, "Částka obdržena:", "L")
	pdf.BottomBlock(width1, "", "L")
	pdf.BottomBlock(width2, "Mezisoučet:", "R")
	pdf.BottomBlock(width3, "3800.00", "R")
	pdf.LineHt(2)

	//-line 2
	pdf.BottomBlock(width1, "Hotově:", "L")
	pdf.BottomBlock(width1, jsn.AmountCash, "L")
	pdf.BottomBlock(width2, "", "R")
	pdf.BottomBlock(width3, "", "R")
	pdf.LineHt(2)

	//-line 3
	pdf.BottomBlock(width1, "Kartou:", "L")
	pdf.BottomBlock(width1, jsn.AmountATM, "L")
	pdf.BottomBlock(width2, "DPH 21 %:", "R")
	pdf.BottomBlock(width3, "798.00", "R")
	pdf.LineHt(2)

	//-line 4
	pdf.BottomBlock(width1, "Převod:", "L")
	pdf.BottomBlock(width1, jsn.AmountOnline, "L")
	pdf.BottomBlock(width2, "", "R")
	pdf.BottomBlock(width3, "", "R")
	pdf.LineHt(2)

	//-line 5
	pdf.BottomBlock(width1, "Dluh:", "L")
	pdf.BottomBlock(width1, jsn.AmountCredit, "L")
	pdf.BottomBlock(width2, "Celková částka:", "R")
	pdf.BottomBlock(width3, "4598.00", "R")

	//*Footer
	pdf.Footer("Rychly servis bohemia 24/7 s.r.o, IČO 17973538, Braunerova 563/7, Libeň, 180 00 Praha 8\nBankovní účet: 5040636073/0800")

	//*Второй лист
	pdf.SecondLeaf("Přejímací protokol:")
	pdf.LineHt(2)
	pdf.SecondLeaf("[✓]  Práce byla převzata bez závad/Shora uvedené zboží bylo na přání namontováno.")
	pdf.LineHt(2)
	pdf.SecondLeaf("[✓]  Faktura je akceptována ohledně ceny a obsahu a položky byly srozumitelně vysvětleny")
	pdf.LineHt(2)
	pdf.SecondLeaf("[✓]  Zaplatim okamžitě bez srážek a nemám žádné")
	pdf.LineHt(10)

	//*Второй лист
	//-Второй блок
	pdf.SecondLeaf("Splatné okamžitě bez srážky")
	pdf.LineHt(2)
	pdf.SecondLeaf("Zadání objednávky/potvrzeni/Povolení")
	pdf.LineHt(2)
	pdf.SecondLeaf("Jsem oprávnen nechat vykonat non zadané práce")
	pdf.LineHt(2)
	pdf.SecondLeaf("Celková účtovaná částka je podle domluvy splatná v hotovosti nebo")
	pdf.LineHt(2)
	pdf.SecondLeaf("platební kartou okamžité na mistě bez srážek. Byl jsem informován o možném")
	pdf.LineHt(2)
	pdf.SecondLeaf("malém poškození a Akcepeji. 2e v pripade malé nedbalosti je ručené vyloučené.")
	pdf.LineHt(2)
	pdf.SecondLeaf("Provedené fakturované položky plati jako dohodmné. Tato ustanovení bylo přečteno a")
	pdf.LineHt(2)

	//*Создаем pdf файл
	err := pdf.OutputFileAndClose()
	if err != nil {
		log.Fatalf("Не удалось вывести файл и закрыть PDF-документ: %v", err)
	}
}
