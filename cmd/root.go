package cmd

import (
	"log"
	"os"

	"github.com/ArturSkrin/kbot/internal/bot"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kbot",
	Short: "kbot — простой Telegram-бот на Golang",
	Long:  "kbot — учебный Telegram-бот, написанный на Golang с использованием cobra и telebot.v4",
	Run: func(cmd *cobra.Command, args []string) {
		runBot()
	},
}

// Execute запускает корневую команду
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func runBot() {
	// Пытаемся подгрузить .env, если он есть (на проде переменные могут быть уже в окружении)
	if err := godotenv.Load(); err != nil {
		log.Println("Файл .env не найден, использую переменные окружения из системы")
	}

	token := os.Getenv("TELE_TOKEN")
	if token == "" {
		log.Fatal("Переменная окружения TELE_TOKEN не задана")
	}

	b, err := bot.NewBot(token)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Бот запущен...")
	b.Start()
}
