package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"log"
	"test/api"
	"test/api/controller"
	"test/api/injection"
	"test/api/middleware"
	"test/api/service"
	"test/api/utils"
)

var (
	db               *gorm.DB                    = injection.CreateDatabase()
	cache            *redis.Client               = injection.SetupRedisConnection()
	Migration        injection.Migration         = injection.NewMigration(db)
	Seed             injection.Seeder            = injection.NewSeeder(db)
	UserService      service.UserService         = service.NewUserService(db, cache)
	BookService      service.BookService         = service.NewBookService(db, cache)
	BorrowService    service.BorrowService       = service.NewBorrowService(db, cache)
	UserController   controller.UserController   = controller.NewUserController(UserService)
	BookController   controller.BookController   = controller.NewBookController(BookService)
	BorrowController controller.BorrowController = controller.NewBorrowController(BorrowService)
	Routes           api.Route                   = api.NewRoute(
		UserController,
		BookController,
		BorrowController,
	)
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

	Routes.Routes(router)

	port := utils.EnvVar("PORT", "8080")
	err := router.Run(":" + port)
	if err != nil {
		log.Println("Failed To Start System")
	}
}
