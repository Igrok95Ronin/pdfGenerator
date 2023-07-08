package main

import (
	"github.com/jung-kurt/gofpdf"
)

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
	heightTable := 16.0
	pdf.CellFormat(widthTable, heightTable, " Popis/Výkon ", "1", 0, "L", true, 0, "")
	pdf.CellFormat(widthTable, heightTable, "Množství", "1", 0, "C", true, 0, "")
	pdf.CellFormat(widthTable, heightTable, "Cena za kus", "1", 0, "C", true, 0, "")
	pdf.MultiCell(widthTable, 8, " Jednotková \n cena bez DPH ", "1", "C", true)

	//*
	//Создаем pdf файл
	err := pdf.OutputFileAndClose("../../yourContract.pdf")
	if err != nil {
		panic(err)
	}
}
