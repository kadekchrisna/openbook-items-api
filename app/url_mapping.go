package app

import (
	"net/http"

	"github.com/kadekchrisna/openbook-items-api/controllers"
)

func mapUrls() {
	router.HandleFunc("/items", controllers.ItemController.Create).Methods(http.MethodPost)
	router.HandleFunc("/items/{id}", controllers.ItemController.Get).Methods(http.MethodGet)
	router.HandleFunc("/items/search", controllers.ItemController.Search).Methods(http.MethodPost)

	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)
}
