package main

import (
	"github.com/solnsumei/simple-chat/config"
	"github.com/solnsumei/simple-chat/models"
)

func runMigrations() {
	config, err := config.LoadConfigVars()

	if err != nil {
		panic("Failed to set config variables")
	}

	err = models.RunMigration(config)
	if err != nil {
		panic(err)
	}
}
