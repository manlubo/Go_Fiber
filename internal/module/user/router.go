package user

import (
	"fiber/internal/module/user/controller"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router, ctl *controller.UserController) {
	r := app.Group("/users")

	r.Post("/", ctl.Create)
	r.Get("/:id", ctl.Get)
	r.Get("/", ctl.List)
	r.Delete("/:id", ctl.Delete)
	r.Patch("/:id", ctl.Update)
}
