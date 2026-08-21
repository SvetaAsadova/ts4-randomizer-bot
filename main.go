package main

import (
	"context"
	"log"

	"simsbot/db"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("8921154285:AAG8bJdbUo33n1s3jbh9F299RmIKk3pHp6c")
	if err != nil {
		log.Panic(err)
	}
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	// Открытие БД
	database, err := db.Open("data/events.db")
	if err != nil {
		log.Printf("⚠️ Не удалось открыть БД: %v", err)
	}
	var repo *db.Repository
	if database != nil {
		repo = db.NewRepository(database)
		defer database.Close()
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		if update.Message.Text == "/start" {
			msg.Text = "Привет! Напиши /event, чтобы получить событие."
		} else if update.Message.Text == "/event" {
			ctx := context.Background()

			if repo == nil {
				msg.Text = "Нет событий для вывода"
			} else {
				hasEvents, err := repo.HasAnyEvents(ctx)
				if err != nil {
					log.Printf("⚠️ Ошибка проверки событий: %v", err)
					msg.Text = "Нет событий для вывода"
				} else if !hasEvents {
					msg.Text = "Нет событий для вывода"
				} else {
					event, err := repo.RandomEvent(ctx)
					if err != nil {
						log.Printf("⚠️ Ошибка чтения события: %v", err)
						msg.Text = "Нет событий для вывода"
					} else if event == nil {
						msg.Text = "Нет событий для вывода"
					} else {
						msg.Text = event.FormatAsCard()
					}
				}
			}
		} else {
			msg.Text = "Я не понимаю эту команду."
		}
		msg.ParseMode = tgbotapi.ModeHTML
		if _, err := bot.Send(msg); err != nil {
			log.Println(err)
		}
	}
}
