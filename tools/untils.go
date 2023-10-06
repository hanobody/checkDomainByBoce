package tools

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func TelegramSender(msg string) {
	bot, err := tgbotapi.NewBotAPI("6632375152:AAHqAthzJfM617m8DRNWTRvwvXJRIWpbd-4")
	if err != nil {
		log.Panic(err)
	}

	// 设置Debug模式，以便查看发送的HTTP请求和响应
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// 替换成你的群组Chat ID
	chatID := int64(-1001779525647)

	// 配置消息
	sendTxt := tgbotapi.NewMessage(chatID, msg)

	// 发送消息
	_, err = bot.Send(sendTxt)
	if err != nil {
		log.Panic(err)
	}
}
