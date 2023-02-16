package main

import (
	"devbook-golang-app/webapp/src/router"
	"devbook-golang-app/webapp/src/utils"
	"fmt"
	"log"
	"net/http"
)

func main() {
	utils.CarregarTemplates()
	r := router.Gerar()

	fmt.Println("Escutando na porte 3000...")
	log.Fatal(http.ListenAndServe(":3000", r))
}
