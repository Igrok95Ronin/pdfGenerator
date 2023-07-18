package pdf

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"io"
	"log"
	"net/http"
	"os"
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
	SecondLeafAT(string)
	Signature(string, string, string, float64, float64, float64, float64, string, string)
	AcceptanceReportAT(float64, string)
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
	p.pdf.SetFont("Arial", "", 10) // Установка шрифта перед выводом текста
	_, lineHt := p.pdf.GetFontSize()
	p.pdf.CellFormat(190, lineHt+20, text, "0", 0, "L", false, 0, "")
}
func (p *pdfDocument) SecondLeafAT(text string) {
	p.pdf.SetFont("Arial", "", 8) // Установка шрифта перед выводом текста
	_, lineHt := p.pdf.GetFontSize()
	p.pdf.MultiCell(190, lineHt*1.5, text, "0", "L", false)

}

// Отчет о приемке
func (p *pdfDocument) AcceptanceReportAT(width float64, text string) {
	p.pdf.SetFont("Arial", "", 8) // Установка шрифта перед выводом текста
	_, lineHt := p.pdf.GetFontSize()
	p.pdf.CellFormat(width, lineHt*1.5, text, "0", 0, "L", false, 0, "")
}

// Подпись
func (p *pdfDocument) Signature(text, alignStr, url string, x, y, w, h float64, webPath, signatureName string) {
	_, lineHt := p.pdf.GetFontSize()
	p.pdf.CellFormat(95, lineHt+20, text, "0", 0, alignStr, false, 0, "")

	//Создание и открытие подписей
	if signatureName == "firstSignature" || signatureName == "secondSignature" {
		resp, err := http.Get("https://app.o95.info/" + webPath)
		if err != nil {
			log.Println(err)
		}
		defer resp.Body.Close()

		out, err := os.Create("../../ui/static/img/" + signatureName + ".jpeg")
		if err != nil {
			log.Println(err)
		}
		defer out.Close()

		_, err = io.Copy(out, resp.Body)
		if err != nil {
			log.Println(err)
		}
	}

	// ImageOptions(src string, x, y, w, h float64, flow bool, options ImageOptions, link int, linkStr string)
	p.pdf.ImageOptions(url, x, y, w, h, false, gofpdf.ImageOptions{ImageType: "jpeg", ReadDpi: true}, 0, "")

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

func Pdf(url string, w http.ResponseWriter) {

	jsn := &json.DataJsonStruct{}
	jsn.Parse(url)

	//PDF
	pdf := newPdfDocument()
	pdf.AddPage()
	pdf.AddUTF8Font() // Добавляем шрифт, поддерживающий больше символов
	pdf.SetFont()     // Задаем шрифт, жирность и размер

	fmt.Println(jsn.Country)
	if jsn.Country == "cz" {

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
			if utf8.RuneCountInString(jsn.Expenses[i].Name) > 25 {
				pdf.TableBody(55.0, 7, jsn.Expenses[i].Name, "L")
			} else {
				pdf.TableBody(55.0, 14, jsn.Expenses[i].Name, "L")
			}
			pdf.TableBody(45.0, 14, strconv.FormatFloat(jsn.Expenses[i].Amount, 'f', -1, 64), "C")
			pdf.TableBody(45.0, 14, strconv.FormatFloat(jsn.Expenses[i].Price, 'f', -1, 64), "C")
			pdf.TableBody(45.0, 14, strconv.FormatFloat(jsn.Expenses[i].PriceBuy*jsn.Expenses[i].Amount, 'f', -1, 64), "C")
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
		pdf.BottomBlock(width3, jsn.Amount, "R")
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
		pdf.BottomBlock(width3, jsn.AmountTax, "R")
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
		pdf.BottomBlock(width3, jsn.AmountTotal, "R")

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
		pdf.LineHt(30)

		//Подпись
		date := jsn.StartedAt
		dateString := date.Format("02 Jan 2006")

		if jsn.SignatureURL != "" {
			pdf.Signature("Datum a podpis objednávky: "+dateString, "L", "../../ui/static/img/firstSignature.jpeg", 10, 230, 70, 0, jsn.SignatureURL, "firstSignature")
		} else {
			pdf.Signature("Datum a podpis objednávky: "+dateString, "L", "../../ui/static/img/defaultSignature.jpeg", 20, 230, 50, 0, jsn.SignatureEndURL, "defaultSignature")
		}
		if jsn.SignatureEndURL != "" {
			pdf.Signature("Datum a podpis objednávky: "+dateString, "R", "../../ui/static/img/secondSignature.jpeg", 125, 230, 70, 0, jsn.SignatureEndURL, "secondSignature")
		} else {
			pdf.Signature("Datum a podpis objednávky: "+dateString, "R", "../../ui/static/img/defaultSignature.jpeg", 135, 230, 50, 0, jsn.SignatureEndURL, "defaultSignature")
		}

		//*Создаем pdf файл
		err := pdf.OutputFileAndClose()
		if err != nil {
			log.Printf("Не удалось вывести файл и закрыть PDF-документ: %v", err)
		}

		// Открываем файл PDF
		pdfFile, err := os.Open("../../yourContract.pdf")
		if err != nil {
			http.Error(w, "Не удалось открыть PDF", http.StatusInternalServerError)
			return
		}
		defer pdfFile.Close()

		// Устанавливаем заголовки для отображения PDF в браузере
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "inline; filename=yourContract.pdf")

		// Пишем содержимое файла в ответ
		_, err = io.Copy(w, pdfFile)
		if err != nil {
			log.Println(err)
		}
	} else if jsn.Country == "at" {

		//*Заголовок документа
		pdf.Header("AT")
		pdf.LineHt(2)

		//*Верхний блок
		//-Первая строка
		pdf.AddText(jsn.Name)

		//-Правый строка ID
		pdf.AddTextRight("Rechnung: " + jsn.RID)

		//-Вторая строка
		pdf.AddText(jsn.Street + " " + jsn.PropertyID)

		//-Третья строка
		pdf.AddText(jsn.ZipCode + " " + jsn.City)
		pdf.LineHt(7)

		//*CheckBox

		if jsn.CheckBox1 == "yes" {
			pdf.AddCheckBox(32.5, "[✓] Bestellung")
		} else {
			pdf.AddCheckBox(32.5, "[ ] Bestellung")
		}
		if jsn.CheckBox2 == "yes" {
			pdf.AddCheckBox(23.5, "[✓] Angebot")
		} else {
			pdf.AddCheckBox(23.5, "[ ] Angebot")
		}
		if jsn.CheckBox3 == "yes" {
			pdf.AddCheckBox(28.5, "[✓] Beratung")
		} else {
			pdf.AddCheckBox(28.5, "[ ] Beratung")
		}
		if jsn.CheckBox4 == "yes" {
			pdf.AddCheckBox(28.5, "[✓] Notdienst")
		} else {
			pdf.AddCheckBox(28.5, "[ ] Notdienst")
		}
		if jsn.CheckBox5 == "yes" {
			pdf.AddCheckBox(23.5, "[✓] Montage")
		} else {
			pdf.AddCheckBox(23.5, "[ ] Montage")
		}
		if jsn.CheckBox6 == "yes" {
			pdf.AddCheckBox(52.5, "[✓] Haushaltsversicherung")
		} else {
			pdf.AddCheckBox(52.5, "[ ] Haushaltsversicherung")
		}
		pdf.LineHt(3)

		//*Таблица
		//-Header
		pdf.TableHeader(55.0, 14, " Bezeichnung / Leistung ", "L")
		pdf.TableHeader(40.0, 14, " Menge ", "C")
		pdf.TableHeader(45.0, 14, " Einzel-Preis (€) ", "C")
		pdf.TableHeader(50.0, 14, " Netto Gesamtpreis (€) ", "C")
		pdf.LineHt(4)

		//-Body
		for i := 0; i < len(jsn.Expenses); i++ {
			if utf8.RuneCountInString(jsn.Expenses[i].Name) > 25 {
				pdf.TableBody(55.0, 7, jsn.Expenses[i].Name, "L")
			} else {
				pdf.TableBody(55.0, 14, jsn.Expenses[i].Name, "L")
			}
			pdf.TableBody(40.0, 14, strconv.FormatFloat(jsn.Expenses[i].Amount, 'f', -1, 64), "C")
			pdf.TableBody(45.0, 14, strconv.FormatFloat(jsn.Expenses[i].Price, 'f', -1, 64), "C")
			pdf.TableBody(50.0, 14, strconv.FormatFloat(jsn.Expenses[i].PriceBuy*jsn.Expenses[i].Amount, 'f', -1, 64), "C")
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
		pdf.BottomBlock(width1, "Betrag erhalten:", "L")
		pdf.BottomBlock(width1, "", "L")
		pdf.BottomBlock(width2, "Nettobetrag:", "R")
		pdf.BottomBlock(width3, jsn.Amount, "R")
		pdf.LineHt(2)

		//-line 2
		pdf.BottomBlock(width1, "Bar:", "L")
		pdf.BottomBlock(width1, jsn.AmountCash, "L")
		pdf.BottomBlock(width2, "", "R")
		pdf.BottomBlock(width3, "", "R")
		pdf.LineHt(2)

		//-line 3
		pdf.BottomBlock(width1, "EC/Kreditkarte:", "L")
		pdf.BottomBlock(width1, jsn.AmountATM, "L")
		pdf.BottomBlock(width2, "MwSt 20 %:", "R")
		pdf.BottomBlock(width3, jsn.AmountTax, "R")
		pdf.LineHt(2)

		//-line 4
		pdf.BottomBlock(width1, "Onlinebanking:", "L")
		pdf.BottomBlock(width1, jsn.AmountOnline, "L")
		pdf.BottomBlock(width2, "", "R")
		pdf.BottomBlock(width3, "", "R")
		pdf.LineHt(2)

		//-line 5
		pdf.BottomBlock(width1, "Per Überweisung:", "L")
		pdf.BottomBlock(width1, jsn.AmountCredit, "L")
		pdf.BottomBlock(width2, "Gesamtsumme:", "R")
		pdf.BottomBlock(width3, jsn.AmountTotal, "R")

		//*Footer
		pdf.Footer("Raiffeisen, IBAN: AT42 3200 0000 1363 8788, BIC: RLNWATWWXXX")

		//*Второй лист
		pdf.SecondLeaf("Auftragserteilung/Bestätigung/Erlaubnis:")
		pdf.LineHt(4.5)
		pdf.SecondLeafAT("Ich bin berechtigt, die von mir in Auftrag gegebenen Arbeiten ausführen zu lassen. Der gesamte Rechnungsbetrag ist in bar oder per EC-Card wie vereinbart sofort vor Ort ohne Abzüge von mir zu entrichten. Auf die Möglichkeit geringfügiger Beschädigung wurde ich hingewiesen und ich akzeptiere, dass für die Öffnungsschäden infolge geringfügiger Fahrlässigkeit die Haftung ausgeschlossen ist. Aufgeführte Rechnungspositionen gelten als fest vereinbart. Regelungen wurden gelesen und bestätigt.")

		//*Второй лист
		//-Второй блок
		pdf.SecondLeaf("Abnahmeprotokoll:")
		pdf.LineHt(5)
		pdf.AcceptanceReportAT(155, " • Wurde die Arbeit ohne Mängel angenommen")
		pdf.AcceptanceReportAT(15, " [✓] Ja")
		pdf.AcceptanceReportAT(20, " [ ] Nein")
		pdf.LineHt(2)
		pdf.AcceptanceReportAT(155, " • Die Rechnung wird in Preis und Inhallt akzeptiert und die Positionen wurden verständlich erklärt")
		pdf.AcceptanceReportAT(15, " [✓] Ja")
		pdf.AcceptanceReportAT(20, " [ ] Nein")
		pdf.LineHt(2)
		pdf.AcceptanceReportAT(155, " • Sind mehrere Produkte angeboten worden")
		pdf.AcceptanceReportAT(15, " [✓] Ja")
		pdf.AcceptanceReportAT(20, " [ ] Nein")
		pdf.LineHt(2)

		//Третий блок
		pdf.SecondLeaf("Die Zahlung werde ich jetzt sofort ohne Abzug vornehmen.")
		pdf.LineHt(2)
		pdf.SecondLeaf(" [✓] keine Beanstandung der Arbeiten und Funktionen")
		pdf.LineHt(2)
		pdf.SecondLeaf(" [✓] für weitere Dienste ist Kontakt erwünscht")
		pdf.LineHt(30)

		//Подпись
		date := jsn.StartedAt
		dateString := date.Format("02 Jan 2006")

		if jsn.SignatureURL != "" {
			pdf.Signature("Datum a podpis objednávky: "+dateString, "L", "../../ui/static/img/firstSignature.jpeg", 10, 230, 70, 0, jsn.SignatureURL, "firstSignature")
		} else {
			pdf.Signature("Datum a podpis objednávky: "+dateString, "L", "../../ui/static/img/defaultSignature.jpeg", 20, 230, 50, 0, jsn.SignatureEndURL, "defaultSignature")
		}
		if jsn.SignatureEndURL != "" {
			pdf.Signature("Datum a podpis objednávky: "+dateString, "R", "../../ui/static/img/secondSignature.jpeg", 125, 230, 70, 0, jsn.SignatureEndURL, "secondSignature")
		} else {
			pdf.Signature("Datum a podpis objednávky: "+dateString, "R", "../../ui/static/img/defaultSignature.jpeg", 135, 230, 50, 0, jsn.SignatureEndURL, "defaultSignature")
		}

		//*Создаем pdf файл
		err := pdf.OutputFileAndClose()
		if err != nil {
			log.Printf("Не удалось вывести файл и закрыть PDF-документ: %v", err)
		}

		// Открываем файл PDF
		pdfFile, err := os.Open("../../yourContract.pdf")
		if err != nil {
			http.Error(w, "Не удалось открыть PDF", http.StatusInternalServerError)
			return
		}
		defer pdfFile.Close()

		// Устанавливаем заголовки для отображения PDF в браузере
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "inline; filename=yourContract.pdf")

		// Пишем содержимое файла в ответ
		_, err = io.Copy(w, pdfFile)
		if err != nil {
			log.Println(err)
		}
	}

}
