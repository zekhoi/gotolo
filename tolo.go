package main

import (
	"net/http"
	"time"
	"tolol/handlers"
	"tolol/utils"
)

var data []byte

func main() {

	go autoUpdate()
	http.HandleFunc("/", handlers.Handler)
	http.ListenAndServe(":8080", nil)
}

func autoUpdate() {
	for range time.Tick(15 * time.Second) {
		utils.UpdateJson()
	}
}
