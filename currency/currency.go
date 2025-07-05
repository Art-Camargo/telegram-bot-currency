package currency

import (
	"encoding/json"
	"fmt"
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
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("erro na resposta da API: status %d", resp.StatusCode)
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

	if coin == "bitcoin" && price < 600_000 {
		telegram.SendMessage(
			fmt.Sprintf("O preço do %s em BRL é: R$ %.2f e ainda não é interessante vender", coin, price),
		)
	} else if coin == "bitcoin" && price >= 600_000 {
		telegram.SendMessage(
			fmt.Sprintf("O preço do %s em BRL é: R$ %.2f e é interessante vender", coin, price),
		)
	} else if coin == "ethereum" && price < 18_000 {
		telegram.SendMessage(
			fmt.Sprintf("O preço do %s em BRL é: R$ %.2f e ainda não é interessante vender", coin, price),
		)
	} else if coin == "ethereum" && price >= 18_000 {
		telegram.SendMessage(
			fmt.Sprintf("O preço do %s em BRL é: R$ %.2f e é interessante vender", coin, price),
		)
	}

}

func RunExchanges() {
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


		sleepOS(1) 
	}
}
