package main

import (
	"net/http"
	"shortly/app/controllers"
	"shortly/app/repositories"
	"shortly/app/services"
)

func main() {
	urlRepository := repositories.NewURLRepository()
	urlService := services.NewURLShortenerService(urlRepository)
	urlController := controllers.NewURLController(urlService)

	http.HandleFunc("/hello", urlController.Index)

	http.ListenAndServe(":5173", nil)
}