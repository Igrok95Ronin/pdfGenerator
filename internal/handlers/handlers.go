package handlers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"net/url"
	"pdgGenerator/internal/pdf"
)

func Pdf(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Начало URL для нашего запроса
	baseURL := "https://app.o95.info/receipt?"

	// Получаем полный URL из запроса
	queryParam := r.URL.String()

	// Парсим URL
	parsedURL, err := url.Parse(queryParam)
	if err != nil {
		log.Println(err)
		return
	}

	// Получаем часть URL после знака '?'
	queryStr := parsedURL.RawQuery

	// Формируем новый URL, объединяя начало baseURL и декодированный URL
	fullURL := fmt.Sprintf("%s%s", baseURL, queryStr)

	// Генерируем PDF, используя новый URL
	pdf.GeneratePdf(fullURL, w)
}
