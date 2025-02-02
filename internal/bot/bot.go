package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log/slog"
	"os"
)

func StartBot() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		slog.Error(err.Error())
	}
	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			handleMessage(bot, update.Message)
		}
	}
}
