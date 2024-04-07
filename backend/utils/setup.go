package utils

import (
	"log/slog"
)

/*
func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
*/

func RunServer(port string) {
	router := GetRouter()
	slog.Info("Starting server", "listenAddr", port)
	err := router.Run()
	if err != nil {
		slog.Error("Oups, something happened. Could not start the server ! ", err)
	}
}
