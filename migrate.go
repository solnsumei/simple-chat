package main

import (
	"github.com/solnsumei/simple-chat/models"
	"github.com/solnsumei/simple-chat/utils"
)

func runMigrations() {
	config, err := utils.LoadConfigVars()

	if err != nil {
		panic("Failed to set config variables")
	}

	err = models.RunMigration(config)
	if err != nil {
		panic(err)
	}
}
