package telegram

func QueryBuilder(telegramBotUrlApi string, telegramChatId string, msg string) string {
	if telegramBotUrlApi == "" || telegramChatId == "" {
		panic("TELEGRAM_URL_API or TELEGRAM_CHAT_ID environment variable is not set")
	}

	query := "?chat_id=" + telegramChatId + "&text=" + msg
	return telegramBotUrlApi + query
}