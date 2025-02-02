package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func handleMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	switch message.Command() {
	case "start":
		msg := tgbotapi.NewMessage(message.Chat.ID, "Привет! Я бот-ассисетент для рассписания") //TODO fix text
		bot.Send(msg)
	case "help":
		msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Команды:\\n/start - Запустить бота\\n/help - Помощь")) //TODO fix text
		bot.Send(msg)
	default:
		msg := tgbotapi.NewMessage(message.Chat.ID, "Неизвестная команда.")
		bot.Send(msg)
	}
}
