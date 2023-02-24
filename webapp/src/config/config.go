package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	//APIUrl representa a URL para comunicação com a API
	APIUrl = "http://localhost"
	//Porta onde a aplicação web está rodando
	Porta = 5000
	// HashKey é utilizada para autenticar o cookie
	HashKey []byte
	// BlockKey é utlizada para criptografar os dados do cookie
	BlockKey []byte
)

//Carregar inicializa as variáveis de ambiente
func Carregar() {
	var erro error

	if erro = godotenv.Load();erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("APP_PORT"))
	if erro != nil {
		log.Fatal(erro)
	}

	APIUrl = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))

}