package telegram

import (
	"log"
	"net/http"
	"os"
)

func SendMessage(msg string) error {
	telegramChatId := os.Getenv("TELEGRAM_CHAT_ID")
	telegramBotUrlApi := os.Getenv("TELEGRAM_URL_API")

	url := QueryBuilder(telegramBotUrlApi, telegramChatId, msg)
	
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		log.Print("Erro ao enviar mensagem para o Telegram:", err)
		return err
	} else {
		log.Print("Mensagem enviada com sucesso para o Telegram:", msg)
	}
	defer resp.Body.Close()

	return nil
}