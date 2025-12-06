//go:build wireinject
// +build wireinject

package board

import (
	"fiber/internal/db"
	"fiber/internal/module/board/controller"
	"fiber/internal/module/board/repository"
	"fiber/internal/module/board/service"

	driver "github.com/arangodb/go-driver"
	"github.com/google/wire"
)

// Board 컬렉션 Provider
func provideBoardCollection(d driver.Database) driver.Collection {
	return db.EnsureCollection(d, "boards")
}

// Wire injector
func InitializeBoardModule() *BoardModule {
	wire.Build(
		db.ConnectArango,
		provideBoardCollection,
		repository.NewBoardRepository,
		service.NewBoardService,
		controller.NewBoardController,
		NewBoardModule,
	)
	return nil
}
