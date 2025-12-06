package board

import (
	"fiber/internal/module/board/controller"

	"github.com/gofiber/fiber/v2"
)

func BoardRouter(app fiber.Router, ctl *controller.BoardController) {
	r := app.Group("/boards")

	r.Post("/", ctl.Create)
	r.Get("/:id", ctl.Get)
	r.Get("/", ctl.List)
	r.Delete("/:id", ctl.Delete)
	r.Patch("/:id", ctl.Update)
}
