package main

import (
	"log"
	"math/rand"

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
	events := []string{
		"Смените работу",
		"Пригласите сима на свидание",
	}
	homeAppliances := []string{
		"Холодильник",
		"Плита",
		"Телевизор",
		"Кофемашина",
	}
	homeApplianceWeights := []int{
		1, // Холодильник — 50% среди бытовой техники
		1, // Плита — 30%
		1, // Телевизор — 10%
		1, // Кофемашина — 20%
	}
	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		if update.Message.Text == "/start" {
			msg.Text = "Привет! Напиши /event, чтобы получить событие."
		} else if update.Message.Text == "/event" {
			totalWeight := 0
			for _, w := range homeApplianceWeights {
				totalWeight += w
			}
			r := rand.Intn(totalWeight)
			sum := 0
			var selectedAppliance string
			for i, w := range homeApplianceWeights {
				sum += w
				if r < sum {
					selectedAppliance = homeAppliances[i]
					break
				}
			}
			if rand.Intn(3) == 0 {
				msg.Text = selectedAppliance + " сломался. Удалите и купите новый"
			} else {
				msg.Text = events[rand.Intn(len(events))]
			}
		} else {
			msg.Text = "Я не понимаю эту команду."
		}
		if _, err := bot.Send(msg); err != nil {
			log.Println(err)
		}
	}
}
