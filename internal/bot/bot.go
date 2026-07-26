package bot

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/telebot.v4"
)

// NewBot создаёт и настраивает объект telebot.Bot
func NewBot(token string) (*telebot.Bot, error) {
	pref := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := telebot.NewBot(pref)
	if err != nil {
		return nil, fmt.Errorf("не удалось создать бота: %w", err)
	}

	registerHandlers(b)
	return b, nil
}

// registerHandlers регистрирует все обработчики сообщений
func registerHandlers(b *telebot.Bot) {
	b.Handle("/start", handleStart)
	b.Handle("/help", handleHelp)
	b.Handle(telebot.OnText, handleText)
}

func handleStart(c telebot.Context) error {
	return c.Send("Привет! Я kbot 🤖\nНапиши /help чтобы узнать что я умею.")
}

func handleHelp(c telebot.Context) error {
	return c.Send("Доступные команды:\n/start — начать работу\n/help — помощь\n\nЛюбой другой текст я просто повторю за тобой (echo).")
}

func handleText(c telebot.Context) error {
	log.Printf("Получено сообщение от %s: %s", c.Sender().Username, c.Text())
	return c.Send("Ты написал: " + c.Text())
}
