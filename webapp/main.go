package main

import (
	"devbook-golang-app/webapp/src/config"
	"devbook-golang-app/webapp/src/router"
	"devbook-golang-app/webapp/src/utils"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()
	utils.CarregarTemplates()
	r := router.Gerar()

	fmt.Printf("Escutando na porta %d\n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
