package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Running API ...")
	config.Load()
	router := router.Gerar()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), router))
}
