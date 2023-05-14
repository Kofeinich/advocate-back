package tgService

import (
	"bot_forge_back/pkg"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func RegNewWebHook(token string, botId string) (err error) {
	webhookURL := pkg.AppConfig.Url + botId
	webhookConfig := tgbotapi.NewWebhook(webhookURL)
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}
	_, err = bot.SetWebhook(webhookConfig)
	if err != nil {
		return err
	}
	bot.ListenForWebhook(webhookURL)
	return nil
}
