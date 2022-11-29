package injection

import (
	"gorm.io/gorm"
	"log"
	"test/api/model"
	"test/api/utils"
)

type Migration interface {
	Migrate()
}

type migrationConnection struct {
	connection *gorm.DB
}

func (m migrationConnection) Migrate() {
	if utils.EnvVar("APP_ENV", "DEVELOPMENT") == "DEVELOPMENT" {
		m.connection.Exec("CREATE DATABASE IF NOT EXISTS")

		errUser := m.connection.Migrator().DropTable(&model.User{})
		if errUser != nil {
			log.Println("Failed To Drop Table User")
		}

		errBook := m.connection.Migrator().DropTable(&model.Book{})
		if errBook != nil {
			log.Println("Failed To Drop Table Book")
		}

		err := m.connection.AutoMigrate(
			&model.User{},
			&model.Book{},
			&model.Borrow{},
		)

		if err != nil {
			log.Println("Failed To Migrate Table")
		}
	}
}

func NewMigration(conn *gorm.DB) Migration {
	return &migrationConnection{
		connection: conn,
	}
}
