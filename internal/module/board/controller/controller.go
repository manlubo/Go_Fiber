package controller

import (
	"fiber/internal/entity"
	"fiber/internal/module/board/service"
	"fiber/internal/util"

	"github.com/gofiber/fiber/v2"
)

type BoardController struct {
	svc *service.BoardService
}

func NewBoardController(svc *service.BoardService) *BoardController {
	return &BoardController{svc: svc}
}

// 생성
func (ctl *BoardController) Create(c *fiber.Ctx) error {
	var board entity.Board
	raw := c.Body()

	if err := util.StrictJSONDecode(raw, &board); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}

	if err := ctl.svc.CreateBoard(&board); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(board)
}

// 조회
func (ctl *BoardController) Get(c *fiber.Ctx) error {
	id := c.Params("id")

	board, err := ctl.svc.GetBoard(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	if board == nil {
		return c.Status(404).JSON(fiber.Map{"error": "not found"})
	}

	return c.JSON(board)
}

// 목록
func (ctl *BoardController) List(c *fiber.Ctx) error {
	boards, err := ctl.svc.GetBoards()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(boards)
}

// 수정
func (ctl *BoardController) Update(c *fiber.Ctx) error {
	raw := c.Body()

	id := c.Params("id")

	var board entity.Board
	if err := util.StrictJSONDecode(raw, &board); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}

	if err := ctl.svc.UpdateBoard(id, &board); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "board updated", "status": "success"})
}

// 삭제
func (ctl *BoardController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := ctl.svc.DeleteBoard(id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "board deleted", "status": "success"})
}
