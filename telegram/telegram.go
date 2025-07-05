package telegram

import (
	"net/http"
	"os"
)

func SendMessage(msg string) error {
	telegramChatId := os.Getenv("TELEGRAM_CHAT_ID")
	telegramBotUrlApi := os.Getenv("TELEGRAM_URL_API")

	url := QueryBuilder(telegramBotUrlApi, telegramChatId, msg)
	
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}