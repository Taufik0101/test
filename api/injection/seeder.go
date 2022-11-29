package injection

import (
	"gorm.io/gorm"
	"test/api/model"
	"test/api/utils"
)

type Seeder interface {
	Faker()
}

type seederConnection struct {
	connection *gorm.DB
}

func (s seederConnection) Faker() {

	if utils.EnvVar("APP_ENV", "DEVELOPMENT") == "DEVELOPMENT" {
		user := model.User{
			Name:  "Admin",
			Email: "admin@gmail.com",
			Phone: "0881523123",
		}

		book := model.Book{
			Name: "Buku Sejarah",
			Page: "300",
		}

		s.connection.Create(&user)
		s.connection.Create(&book)
	}
}

func NewSeeder(conn *gorm.DB) Seeder {
	return &seederConnection{connection: conn}
}
