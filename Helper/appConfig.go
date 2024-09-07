package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"productapi/models"
)

func LoadConfigs() models.AppConfig {
	var config models.AppConfig

	file, err := os.Open("appsettings.json")

	if err != nil {
		fmt.Println("Can't open file")
		return config
	}

	defer file.Close()

	bytes, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Println("Can't read")

		return config
	}

	err = json.Unmarshal(bytes, &config)

	if err != nil {
		fmt.Println("Can't serialize ", err)

		return config
	}

	return config

}
