package botService

type Service struct {
	r repo
}

type repo interface {
	DeleteBotFromList(botID string) error
	GelAllBotsFromList() ([]string, error)
	AddBotToBotsList(botID string) error
	CreateBotConfig(botID string, config string) error
	GetBotConfigByID(botID string) (string, error)
	CreateBotToken(botID string, token string) error
	GetBotTokenByID(botID string) (string, error)
}

func NewService(r repo) *Service {
	return &Service{r: r}
}
