package databaselayer

import (
	"fmt"
	"productapi/helper"
)

func GetConnection() {

	var config helper.AppConfig = helper.LoadConfigs()

	fmt.Printf("Port: %d\n", config.Database.Port)
}
