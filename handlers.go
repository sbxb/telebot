package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleCommand(um *tgbotapi.Message) tgbotapi.Chattable {
	msg := tgbotapi.NewMessage(um.Chat.ID, "")
	msg.ParseMode = "html"
	msg.Text = "That was a <u>command</u>, wasn't that\n"
	msg.Text += "It was <b>" + um.Command() + "</b>\n"
	msg.Text += "Arguments were <i>" + um.CommandArguments() + "</i>\n"

	return msg
}

func handleText(um *tgbotapi.Message) tgbotapi.Chattable {
	msg := tgbotapi.NewMessage(um.Chat.ID, um.Text)
	msg.ReplyToMessageID = um.MessageID

	return msg
}
