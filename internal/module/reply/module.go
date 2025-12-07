package reply

import (
	"fiber/internal/module/reply/controller"

	driver "github.com/arangodb/go-driver"
	"github.com/gofiber/fiber/v2"
)

type ReplyModule struct {
	db  driver.Database
	ctl *controller.ReplyController
}

// 모듈 생성자
func NewReplyModule(db driver.Database, ctl *controller.ReplyController) *ReplyModule {
	return &ReplyModule{
		db:  db,
		ctl: ctl,
	}
}

// 모듈 초기화 + 라우팅 등록
func (m *ReplyModule) Register(app *fiber.App) {
	// 라우터 등록
	ReplyRouter(app, m.ctl)
}
