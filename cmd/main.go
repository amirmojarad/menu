package main

import (
	"menu_generator/cmd/server"
	"menu_generator/internal/logger"
)

func main() {
	if err := server.RunServer(); err != nil {
		logger.GetLogger().Fatal(err)

		return
	}

}
