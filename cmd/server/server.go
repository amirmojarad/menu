package server

import "menu_generator/config"

func RunServer() error {
	cfg, err := config.GetConfig()
	if err != nil {
		return err
	}

	return nil
}
