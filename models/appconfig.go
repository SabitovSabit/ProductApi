package models


type AppConfig struct {
	Database struct {
		Server   string
		Port     int
		Database string
	}
}