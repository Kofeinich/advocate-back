package botService

func (s Service) UpdateBotConfig(conf string, id string) error {
	err := s.r.CreateBotConfig(id, conf)
	if err != nil {
		return err
	}
	return nil
}
