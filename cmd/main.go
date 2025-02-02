package main

import (
	"fmt"
	"github.com/BioMihanoid/reminder_schedule_bot/config"
	"github.com/BioMihanoid/reminder_schedule_bot/internal/bot"
)

func main() {
	fmt.Printf("Start tg bot...\n")
	config.LoadConfig()
	bot.StartBot()
}
