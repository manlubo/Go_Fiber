package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Cors() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000, http://localhost:5173",
		AllowHeaders:     "*",
		AllowMethods:     "*",
		AllowCredentials: true,
	})
}
