package main

import (
	"DevBook/webapp/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Rodando webapp...")

	r := router.Gerar()
	log.Fatal(http.ListenAndServe(":3000", r))
}
