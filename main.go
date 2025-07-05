package main

import (
	"log"

	"github.com/Art-Camargo/currency-manager/currency"
	"github.com/joho/godotenv"
)


func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar .env:", err)
	}

	currency.RunExchanges()
}