package botService

func (s Service) DeleteBot(id string) error {
	err := s.r.DeleteBotFromList(id)
	if err != nil {
		return err
	}
	return nil
}
