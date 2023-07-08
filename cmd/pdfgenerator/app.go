package main

import (
	"github.com/jung-kurt/gofpdf"
)

// test
func main() {
	pdf := gofpdf.New("P", "mm", "A4", "")                               //
	pdf.AddPage()                                                        //
	pdf.SetFont("Arial", "B", 16)                                        // Задаем шрифт, жирность и размер
	pdf.AddUTF8Font("Arial", "", "../../ui/static/fonts/DejaVuSans.ttf") // Добавляем шрифт, поддерживающий больше символов
	_, lineHt := pdf.GetFontSize()                                       // высота строки зависит от размера шрифта

	//*
	//Заголовок документа
	pdf.CellFormat(190, lineHt, "F A K T U R A", "0", 0, "C", false, 0, "")

	//*
	//Верхний блок
	//Первая строка
	pdf.SetFont("Arial", "", 14)
	pdf.Ln(lineHt * 2) //переход на новую строку
	pdf.CellFormat(95, lineHt, "Kamil Teplý", "0", 0, "L", false, 0, "")

	//*
	//Правый блок ID
	pdf.SetFont("Arial", "B", 14)
	pdf.SetTextColor(52, 116, 178) //задает цвет текста rgb
	pdf.CellFormat(95, lineHt, "ID: CZ-3135", "0", 0, "R", false, 0, "")

	//*
	//Вторая строка
	pdf.SetFont("Arial", "", 14)
	pdf.SetTextColor(0, 0, 0)
	pdf.Ln(lineHt)
	pdf.CellFormat(190, lineHt, "Francouzská 2", "0", 0, "L", false, 0, "")

	//Третья строка
	pdf.Ln(lineHt)
	pdf.CellFormat(190, lineHt, "12000 Praha", "0", 0, "L", false, 0, "")

	//Radio кнопки
	pdf.SetFont("Arial", "", 10)
	pdf.Ln(lineHt * 5)
	pdf.CellFormat(38, lineHt, "[✓] Objednavka", "0", 0, "C", false, 0, "")
	pdf.CellFormat(25, lineHt, "[ ] Nabidka", "0", 0, "C", false, 0, "")
	pdf.CellFormat(37, lineHt, "[ ] Konzultace", "0", 0, "C", false, 0, "")
	pdf.CellFormat(37, lineHt, "[✓] Nalehavost", "0", 0, "C", false, 0, "")
	pdf.CellFormat(25, lineHt, "[ ] Montaz", "0", 0, "C", false, 0, "")
	pdf.CellFormat(31.6, lineHt, "[✓] Pojisteni", "0", 0, "C", false, 0, "")

	//*Таблица
	//Header
	pdf.Ln(10)
	pdf.SetFillColor(52, 116, 178)  // Установка цвета заливки для заголовка
	pdf.SetTextColor(255, 255, 255) // Устанавливает цвет текста
	pdf.SetDrawColor(227, 227, 227) // Устанавливаем цвет границы в синий

	widthTable := 48.5
	heightTable := 14.0
	pdf.CellFormat(widthTable, heightTable, " Popis/Výkon ", "1", 0, "L", true, 0, "")
	pdf.CellFormat(widthTable, heightTable, "Množství", "1", 0, "C", true, 0, "")
	pdf.CellFormat(widthTable, heightTable, "Cena za kus", "1", 0, "C", true, 0, "")
	pdf.MultiCell(widthTable, 7, " Jednotková \n cena bez DPH ", "1", "C", true)

	//Body
	pdf.Ln(0)
	pdf.SetFillColor(255, 255, 255) // Установка цвета заливки для заголовка
	pdf.SetTextColor(0, 0, 0)       // Устанавливает цвет текста

	pdf.CellFormat(widthTable, heightTable, " Vyjezd ", "1", 0, "L", true, 0, "")
	pdf.CellFormat(widthTable, heightTable, "2", "1", 0, "C", true, 0, "")
	pdf.CellFormat(widthTable, heightTable, "600.00", "1", 0, "C", true, 0, "")
	pdf.CellFormat(widthTable, heightTable, "1200.00", "1", 0, "C", true, 0, "")

	pdf.Ln(14)
	pdf.CellFormat(widthTable, heightTable, " Odmitani ", "1", 0, "L", true, 0, "")
	pdf.CellFormat(widthTable, heightTable, "1", "1", 0, "C", true, 0, "")
	pdf.CellFormat(widthTable, heightTable, "2600.00", "1", 0, "C", true, 0, "")
	pdf.CellFormat(widthTable, heightTable, "2600.00", "1", 0, "C", true, 0, "")

	//*Нижний блок
	//line 1
	pdf.Ln(20)
	widthDownBlock := 48.5
	pdf.CellFormat(widthDownBlock, lineHt+2, " Částka obdržena: ", "0", 0, "L", false, 0, "")
	pdf.CellFormat(widthDownBlock, lineHt+2, "  ", "0", 0, "C", false, 0, "")
	pdf.CellFormat(widthDownBlock+20.5, lineHt+2, " Mezisoučet: ", "0", 0, "R", false, 0, "")
	pdf.CellFormat(28, lineHt+2, " 3800.00 ", "0", 0, "R", false, 0, "")
	//line 2
	pdf.Ln(lineHt + 2)
	pdf.CellFormat(widthDownBlock, lineHt+2, " Hotově: ", "0", 0, "L", false, 0, "")
	pdf.CellFormat(widthDownBlock, lineHt+2, " 0.00 ", "0", 0, "C", false, 0, "")
	pdf.CellFormat(widthDownBlock+20.5, lineHt+2, "  ", "0", 0, "R", false, 0, "")
	pdf.CellFormat(28, lineHt+2, "  ", "0", 0, "R", false, 0, "")
	//line 3
	pdf.Ln(lineHt + 2)
	pdf.CellFormat(widthDownBlock, lineHt+2, " Kartou: ", "0", 0, "L", false, 0, "")
	pdf.CellFormat(widthDownBlock, lineHt+2, " 0.00 ", "0", 0, "C", false, 0, "")
	pdf.CellFormat(widthDownBlock+20.5, lineHt+2, " DPH 21 %: ", "0", 0, "R", false, 0, "")
	pdf.CellFormat(28, lineHt+2, " 798.00 ", "0", 0, "R", false, 0, "")
	//line 4
	pdf.Ln(lineHt + 2)
	pdf.CellFormat(widthDownBlock, lineHt+2, " Převod: ", "0", 0, "L", false, 0, "")
	pdf.CellFormat(widthDownBlock, lineHt+2, " 0.00 ", "0", 0, "C", false, 0, "")
	pdf.CellFormat(widthDownBlock+20.5, lineHt+2, "  ", "0", 0, "R", false, 0, "")
	pdf.CellFormat(28, lineHt+2, "  ", "0", 0, "R", false, 0, "")
	//line 5
	pdf.Ln(lineHt + 2)
	pdf.CellFormat(widthDownBlock, lineHt+2, " Dluh: ", "0", 0, "L", false, 0, "")
	pdf.CellFormat(widthDownBlock, lineHt+2, " 4598.00 ", "0", 0, "C", false, 0, "")
	pdf.CellFormat(widthDownBlock+20.5, lineHt+2, " Celková částka: ", "0", 0, "R", false, 0, "")
	pdf.CellFormat(28, lineHt+2, " 4598.00 ", "0", 0, "R", false, 0, "")

	//*Нижний блок
	// Задаем размер нижнего поля
	bottomMargin := 26.0
	// Получаем размеры страницы
	_, pageHeight := pdf.GetPageSize()
	// Устанавливаем новую высоту Y, вычитая нижний отступ и высоту строки из высоты страницы
	pdf.SetY(pageHeight - bottomMargin - lineHt)
	pdf.SetFont("Arial", "", 11) // Установка шрифта перед выводом текста
	pdf.MultiCell(190, lineHt, "Rychly servis bohemia 24/7 s.r.o, IČO 17973538, Braunerova 563/7, Libeň, 180 00 Praha 8\nBankovní účet: 5040636073/0800", "0", "C", false)

	//*Создаем pdf файл
	err := pdf.OutputFileAndClose("../../yourContract.pdf")
	if err != nil {
		panic(err)
	}
}
