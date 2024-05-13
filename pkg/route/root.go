package route

import (
	"github.com/gofiber/fiber/v2"
	swagger "github.com/gofiber/swagger"
)

func GeneralRoute(a *fiber.App) {
	a.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to Go Auth API!",
			"docs":    "/swagger/index.html",
			"status":  "/h34l7h",
		})
	})

}

func SwaggerRoute(a *fiber.App) {
	// Create route group.
	route := a.Group("/swagger")
	route.Get("*", swagger.HandlerDefault)
}

func NotFoundRoute(a *fiber.App) {
	a.Use(
		func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "sorry, endpoint is not found",
			})
		},
	)
}
