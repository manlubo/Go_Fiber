package controller

import (
	"fiber/internal/entity"
	"fiber/internal/module/user/service"
	"fiber/internal/util"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	svc *service.UserService
}

func NewUserController(svc *service.UserService) *UserController {
	return &UserController{svc: svc}
}

func (ctl *UserController) Create(c *fiber.Ctx) error {
	var user entity.User
	raw := c.Body()

	if err := util.StrictJSONDecode(raw, &user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}

	if err := ctl.svc.CreateUser(&user); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(user)
}

func (ctl *UserController) Get(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := ctl.svc.GetUser(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	if user == nil {
		return c.Status(404).JSON(fiber.Map{"error": "not found"})
	}

	return c.JSON(user)
}

func (ctl *UserController) List(c *fiber.Ctx) error {
	users, err := ctl.svc.GetUsers()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(users)
}

func (ctl *UserController) Update(c *fiber.Ctx) error {
	raw := c.Body()

	id := c.Params("id")

	var user entity.User
	if err := util.StrictJSONDecode(raw, &user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}

	if err := ctl.svc.UpdateUser(id, &user); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "user updated", "status": "success"})
}

func (ctl *UserController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := ctl.svc.DeleteUser(id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "user deleted", "status": "success"})
}
