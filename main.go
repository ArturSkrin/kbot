package main

import (
	"log"
	"os"

	"github.com/ArturSkrin/kbot/cmd"
)

// TELE_TOKEN берётся из переменной окружения / .env
// Обработчики используют telebot.Context для ответа на сообщения
func main() {
	if os.Getenv("TELE_TOKEN") == "" {
		log.Println("TELE_TOKEN не задан в окружении")
	}
	cmd.Execute()
}
