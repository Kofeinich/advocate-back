package repository

func BotConfigKey(id string) string {
	return "botConfig_" + id
}

func BotTokenKey(id string) string {
	return "botToken_" + id
}

func BotUserStateKey(botID string, userID string) string {
	return botID + "_" + userID
}

const AllBotsKey = "allBots"
