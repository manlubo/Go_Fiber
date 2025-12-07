package controller

import (
	"fiber/internal/entity"
	"fiber/internal/module/reply/service"
	"fiber/internal/util"

	"github.com/gofiber/fiber/v2"
)

type ReplyController struct {
	svc *service.ReplyService
}

func NewReplyController(svc *service.ReplyService) *ReplyController {
	return &ReplyController{svc: svc}
}

func (ctl *ReplyController) Create(c *fiber.Ctx) error {
	var reply entity.Reply
	raw := c.Body()

	if err := util.StrictJSONDecode(raw, &reply); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}

	if err := ctl.svc.CreateReply(&reply); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(reply)
}

// 조회
func (ctl *ReplyController) Get(c *fiber.Ctx) error {
	id := c.Params("id")

	reply, err := ctl.svc.GetReply(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	if reply == nil {
		return c.Status(404).JSON(fiber.Map{"error": "not found"})
	}

	return c.JSON(reply)
}

// 목록
func (ctl *ReplyController) List(c *fiber.Ctx) error {
	replies, err := ctl.svc.GetReplies()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(replies)
}

// 수정
func (ctl *ReplyController) Update(c *fiber.Ctx) error {
	raw := c.Body()

	id := c.Params("id")

	var reply entity.Reply
	if err := util.StrictJSONDecode(raw, &reply); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}

	if err := ctl.svc.UpdateReply(id, &reply); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "reply updated", "status": "success"})
}

// 삭제
func (ctl *ReplyController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := ctl.svc.DeleteReply(id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "reply deleted", "status": "success"})
}
