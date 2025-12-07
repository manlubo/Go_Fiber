package reply

import (
	"fiber/internal/module/reply/controller"

	"github.com/gofiber/fiber/v2"
)

func ReplyRouter(app fiber.Router, ctl *controller.ReplyController) {
	r := app.Group("/replies")

	r.Post("/", ctl.Create)
	r.Get("/:id", ctl.Get)
	r.Get("/", ctl.List)
	r.Patch("/:id", ctl.Update)
	r.Delete("/:id", ctl.Delete)
}
