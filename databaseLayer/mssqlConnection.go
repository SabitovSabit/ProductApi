package databaselayer

import (
	"fmt"

	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"

	"github.com/labstack/gommon/log"

	"productapi/helper"
	"productapi/models"
)

func GetConnection() *sql.DB {

	var config models.AppConfig = helper.LoadConfigs()

	connString := fmt.Sprintf("server=%s;port=%d;database=%s;integrated security=true;", config.Database.Server, config.Database.Port, config.Database.Database)

	db, err := sql.Open("sqlserver", connString)

	if err != nil {
		log.Error("Unable to connect sql server %v\n", err)
	}

	return db
}
