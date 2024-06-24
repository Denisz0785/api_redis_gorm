package router

import (
	"github.com/gofiber/fiber/v2"

	"redis_gorm_fiber/controller"
)

// NewRouter initializes the Fiber app and sets up the routes.
func NewRouter(router *fiber.App, novelController *controller.NovelController) *fiber.App {
	// Home route
	router.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, World!")
	})

	// Novel creation route
	router.Post("/novel", novelController.CreateNovel)

	// Get novel by ID route
	router.Get("/novel/:id", novelController.GetNovelById)

	// Route to delete a novel by ID
	router.Delete("/novel/:id", novelController.DeleteNovel)

	// Route to update a novel by ID
	router.Put("/novel/:id", novelController.UpdateNovel)

	return router
}
