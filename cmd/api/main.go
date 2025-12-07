package main

import (
	"github.com/gofiber/fiber/v2"

	"fiber/internal/module/board"
	"fiber/internal/module/reply"
	"fiber/internal/module/user"
)

func main() {
	app := fiber.New()

	// 모듈 등록
	// user
	userModule := user.InitializeUserModule()
	userModule.Register(app)

	// board
	boardModule := board.InitializeBoardModule()
	boardModule.Register(app)

	// reply
	replyModule := reply.InitializeReplyModule()
	replyModule.Register(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]interface{}{"message": "Hello, World!", "status": "success"})
	})

	app.Listen(":3000")
}
