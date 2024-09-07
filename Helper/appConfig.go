package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type AppConfig struct {
	Database struct {
		Server   string
		Port     int
		Database string
	}
}

func LoadConfigs() AppConfig {
	var config AppConfig

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
