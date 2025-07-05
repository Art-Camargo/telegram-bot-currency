package currency

import "os"

func QueryBuilder(cryptoId string, vsCurrency string) string {
	baseUrl := os.Getenv("COINGECKO_API_URL")
	if baseUrl == ""  {
		panic("COINGECKO_API_URL environment variable is not set")
	}

	query := "?ids=" + cryptoId + "&vs_currencies=" + vsCurrency
	return baseUrl + query
}