package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	token := os.Getenv("TOKEN")
	if token == "" {
		log.Fatalln("Error: no token provided")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		var msg tgbotapi.Chattable

		if update.Message.IsCommand() {
			msg = handleCommand(update.Message)
		} else {
			msg = handleText(update.Message)
		}

		bot.Send(msg)
	}
}
