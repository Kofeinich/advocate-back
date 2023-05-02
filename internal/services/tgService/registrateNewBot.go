package tgService

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func RegNewBot(token string, botId string) (err error) {
	webhookURL := "https://asporto.serveo.net/tg_webhook/" + botId
	webhookConfig := tgbotapi.NewWebhook(webhookURL)
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}
	_, err = bot.SetWebhook(webhookConfig)
	if err != nil {
		return err
	}
	return nil
}
