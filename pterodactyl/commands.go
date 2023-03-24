package pterodactyl

func SendCommand(command string, serverID string) error {
	s, err := GetServer(serverID)
	if err != nil {
		return err
	}
	if err := s.SendCommand(command); err != nil {
		return err
	}

	return nil
}
