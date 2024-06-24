package main

import (
	"log"
	"redis_gorm_fiber/config"
	"redis_gorm_fiber/controller"
	"redis_gorm_fiber/database"
	"redis_gorm_fiber/model"
	"redis_gorm_fiber/repo"
	"redis_gorm_fiber/router"
	"redis_gorm_fiber/usecase"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// main is the entry point of the application.
// It loads the configuration, establishes connections to MySQL and Redis,
// and starts the server.
func main() {
	// Load the application configuration
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load config", err)
	}

	// Establish a connection to the MySQL database
	db := database.ConnectionMySQLDB(&loadConfig)

	// Establish a connection to the Redis database
	rdb := database.ConnectionRedisDB(&loadConfig)

	// Auto-migrate the `Novel` model to create the corresponding table in MySQL
	db.AutoMigrate(&model.Novel{})

	// Start the server
	startServer(db, rdb)
}

// startServer initializes the server and starts it.
func startServer(db *gorm.DB, rdb *redis.Client) {
	// Create a new Fiber instance
	app := fiber.New()

	// Create a new NovelRepo instance with the given database and Redis connection
	novelRepo := repo.NewNovelRepo(db, rdb)

	// Create a new NovelUseCase instance with the NovelRepo
	novelUseCase := usecase.NewNovelUseCase(novelRepo)

	// Create a new NovelController instance with the NovelUseCase
	novelController := controller.NewNovelController(novelUseCase)

	// Create a new router with the Fiber app and NovelController
	routes := router.NewRouter(app, novelController)

	// Start the server on port 3400
	err := routes.Listen(":3400")

	// If there is an error starting the server, log the error and exit
	if err != nil {
		log.Fatal(err)
	}
}
