package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"test/api/injection"
	"test/api/middleware"
	"test/api/utils"
)

var (
	db        *gorm.DB            = injection.CreateDatabase()
	Migration injection.Migration = injection.NewMigration(db)
	Seed      injection.Seeder    = injection.NewSeeder(db)
)

func main() {
	defer injection.CloseDatabaseConnection(db)

	Migration.Migrate()
	Seed.Faker()

	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	port := utils.EnvVar("PORT", "8080")
	err := router.Run(":" + port)
	if err != nil {
		log.Println("Failed To Start System")
	}
}
