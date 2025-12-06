package board

import (
	"fiber/internal/module/board/controller"

	driver "github.com/arangodb/go-driver"
	"github.com/gofiber/fiber/v2"
)

type BoardModule struct {
	db  driver.Database
	ctl *controller.BoardController
}

// 모듈 생성자
func NewBoardModule(db driver.Database, ctl *controller.BoardController) *BoardModule {
	return &BoardModule{
		db:  db,
		ctl: ctl,
	}
}

// 모듈 초기화 + 라우팅 등록
func (m *BoardModule) Register(app *fiber.App) {
	// 라우터 등록
	BoardRouter(app, m.ctl)
}
