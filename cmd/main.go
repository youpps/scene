package main

import (
	"fmt"
	"log"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"githuub.com/youpps/scenes"
)

var channelId int64 = 0
var groupId int64 = 0

func main() {
	bot, err := tg.NewBotAPI("")
	if err != nil {
		log.Fatalln(err)
	}

	startScene := scenes.NewScene("start")

	startScene.AddJoinHandler(func(update tg.Update) {
		chatId := update.Message.Chat.ID
		msg := tg.NewMessage(chatId, "JJJ")
		bot.Send(msg)
	})

	startScene.AddNextHandler(func(updateData *scenes.UpdateData) {
		if updateData.Update.Message.ForwardFromChat != nil {
			channelId = updateData.Update.Message.ForwardFromChat.ID
			fmt.Println(channelId)
			updateData.Ctx.Next()
		}
	})

	startScene.AddNextHandler(func(updateData *scenes.UpdateData) {
		if updateData.Update.Message != nil {
			groupId = updateData.Update.Message.Chat.ID
			fmt.Println(channelId)
			updateData.Ctx.Leave()
		}
	})

	scenesManager := scenes.NewScenesManager(startScene)

	updatesConfig := tg.NewUpdate(0)
	updatesConfig.Timeout = 60
	updates := bot.GetUpdatesChan(updatesConfig)

	for update := range updates {
		scenesManager.Handle(update, func(update tg.Update) {
			if update.Message != nil && update.Message.Text == "/start" {
				scenesManager.Join(update, startScene)
				return
			}
			fmt.Println(groupId, channelId)
			if update.Message != nil && groupId != 0 && channelId != 0 {
				msg := tg.NewMessage(groupId, "sadas")
				bot.Send(msg)
			}
		})
	}
}
