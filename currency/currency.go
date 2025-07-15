package currency

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"time"

	"github.com/Art-Camargo/currency-manager/telegram"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func sleepOS(totalInMinutes int) {
	time.Sleep(time.Duration(totalInMinutes) * time.Minute)
}

func execExchange(coin string, currency string) (float32, error) {
	resp, err := http.Get(QueryBuilder(coin, currency))
	if err != nil {
		return 0, fmt.Errorf("erro ao fazer requisição: %v", err)
	} else {
		log.Printf("Requisição feita com sucesso para %s em %s", coin, currency)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("erro na resposta da API: status %d", resp.StatusCode)
	} else {
		log.Printf("Resposta da API recebida com sucesso: status %d", resp.StatusCode)
	}

	var data map[string]map[string]float32
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, fmt.Errorf("erro ao decodificar resposta: %v", err)
	}

	price := data[coin][strings.ToLower(currency)]
	return price, nil
}

func sendTelegramMessage(coin string, price float32) {
  if(coin == "" || price <= 0) {
		return
	}
	coin = strings.ToLower(coin)
	
	if coin == "bitcoin" {
		if price < 530_000 {
			telegram.SendMessage(
				fmt.Sprintf("O preço do %s em BRL é: R$ %.2f. CHAME IMEDIATAMENTE O ARTUR, POIS A COMPRA ESTÁ INTERESSANTE", coin, price),
			)
		} 
	} else if coin == "ethereum" {
		if price < 8_000 {
			telegram.SendMessage(
				fmt.Sprintf("O preço do %s em BRL é: R$ %.2f. CHAME IMEDIATAMENTE O ARTUR, POIS A COMPRA ESTÁ INTERESSANTE", coin, price),
			)
		} else if price > 20_000 {
			telegram.SendMessage(
				fmt.Sprintf("O preço do %s em BRL é: R$ %.2f. CHAME IMEDIATAMENTE O ARTUR, POIS A VENDA ESTÁ INTERESSANTE", coin, price),
			)
		}
	}
}

func RunExchanges() {
	log.Print("Iniciando o bot de monitoramento de moedas...")
	coins := []string{"bitcoin", "ethereum"}

	for {
		for _, coin := range coins {
			price, err := execExchange(coin, "BRL")
			if err != nil {
				fmt.Printf("Erro ao buscar preço de %s: %v\n", coin, err)
				continue
			}

			
			title := cases.Title(language.Portuguese)
			capitalizedCoin := title.String(coin)
		
			sendTelegramMessage(capitalizedCoin, price)
		}


		sleepOS(5) 
	}
}
