package main

import (
	"devbook-golang-app/api/src/config"
	"devbook-golang-app/api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()
	r := router.Gerar()
	fmt.Printf("Escutando na porta: %d", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
