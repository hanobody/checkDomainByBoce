package tools

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/spf13/viper"
)

func TelegramSender(msg string) {
	bot, err := tgbotapi.NewBotAPI(viper.GetString("telegram.bottoken"))
	if err != nil {
		log.Panic(err)
	}

	// 设置Debug模式，以便查看发送的HTTP请求和响应
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// 替换成你的群组Chat ID
	chatID := int64(viper.GetInt("telegram.groupid"))

	// 配置消息
	sendTxt := tgbotapi.NewMessage(chatID, msg)

	// 发送消息
	_, err = bot.Send(sendTxt)
	if err != nil {
		log.Panic(err)
	}
}
func InitConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("config app inited")
}
