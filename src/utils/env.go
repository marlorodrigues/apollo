package utils

import (
	"github.com/joho/godotenv"
)

func LoadConfig() string {
	err := godotenv.Load()
	if err != nil {
		logger.Errorln("Erro ao carregar o arquivo .env")
		logger.Errorln("Buscando como variavel de sistema...")
	}

	return ""
}
