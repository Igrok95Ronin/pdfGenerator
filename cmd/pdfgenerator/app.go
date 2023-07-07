package main

import (
	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "") //
	pdf.AddPage()                          //
	pdf.SetFont("Arial", "B", 16)          // Задаем шрифт, жирность и размер

	//Заголовок документа
	_, lineHt := pdf.GetFontSize() // высота строки зависит от размера шрифта
	pdf.CellFormat(190, lineHt, "F A K T U R A", "0", 0, "C", false, 0, "")

	//Левый верхний блок
	pdf.SetFont("Arial", "", 12)
	pdf.Ln(lineHt * 2) //переход на новую строку
	pdf.CellFormat(95, lineHt, "Kamil Teplý", "0", 0, "L", false, 0, "")

	pdf.SetFont("Arial", "B", 12)
	pdf.SetTextColor(52, 116, 178) //задает цвет текста rgb
	pdf.CellFormat(95, lineHt, "ID: CZ-3135", "0", 0, "R", false, 0, "")

	pdf.SetFont("Arial", "", 12)
	pdf.SetTextColor(0, 0, 0)
	pdf.Ln(lineHt)
	pdf.CellFormat(190, lineHt, "Francouzská 2", "0", 0, "L", false, 0, "")

	//
	err := pdf.OutputFileAndClose("yourContract.pdf")
	if err != nil {
		panic(err)
	}
}
