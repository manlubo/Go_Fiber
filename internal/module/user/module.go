package user

import (
	"fiber/internal/module/user/controller"

	driver "github.com/arangodb/go-driver"
	"github.com/gofiber/fiber/v2"
)

type UserModule struct {
	db  driver.Database
	ctl *controller.UserController
}

// 모듈 생성자
func NewUserModule(db driver.Database, ctl *controller.UserController) *UserModule {
	return &UserModule{
		db:  db,
		ctl: ctl,
	}
}

// 모듈 초기화 + 라우팅 등록
func (m *UserModule) Register(app *fiber.App) {
	// 라우터 등록
	UserRouter(app, m.ctl)
}
