package route

import (
	"github.com/amirrstm/go-clinique/app/controller"
	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public route.
func PublicRoutes(a *fiber.App) {
	// Create route group.
	route := a.Group("/api/v1")

	route.Post("/login", controller.GetNewAccessToken)

}
