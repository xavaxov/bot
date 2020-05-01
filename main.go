package main

import (
	"bot/answers"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("1190159727:AAFZ6dkY9OQtVY0zCYD6MkEpXD8wvDUiwJg")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		text := answers.Answer(update.Message.Text) + update.Message.Chat.UserName
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}

// func answer(a string) string {
// 	switch a {
// 	case "❤️":
// 		return "Love you too, "
// 	default:
// 		return "Send me love ❤️, "
// 	}
// }
