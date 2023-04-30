package botService

func (s Service) GetAllBots() ([]string, error) {
	bots, err := s.r.GelAllBotsFromList()
	if err != nil {
		return nil, err
	}
	return bots, nil
}
