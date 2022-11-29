package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"log"
	"test/api/injection"
	"test/api/middleware"
	"test/api/utils"
)

var (
	db        *gorm.DB            = injection.CreateDatabase()
	cache     *redis.Client       = injection.SetupRedisConnection()
	Migration injection.Migration = injection.NewMigration(db)
	Seed      injection.Seeder    = injection.NewSeeder(db)
)

func main() {
	defer injection.CloseDatabaseConnection(db)
	defer func(cache *redis.Client) {
		err := cache.Close()
		if err != nil {

		}
	}(cache)

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
