//go:build wireinject
// +build wireinject

package reply

import (
	"fiber/internal/db"
	"fiber/internal/module/reply/controller"
	"fiber/internal/module/reply/repository"
	"fiber/internal/module/reply/service"

	driver "github.com/arangodb/go-driver"
	"github.com/google/wire"
)

// Reply 컬렉션 Provider
func provideReplyCollection(d driver.Database) driver.Collection {
	return db.EnsureCollection(d, "replies")
}

// Wire injector
func InitializeReplyModule() *ReplyModule {
	wire.Build(
		db.ConnectArango,
		provideReplyCollection,
		repository.NewReplyRepository,
		service.NewReplyService,
		controller.NewReplyController,
		NewReplyModule,
	)
	return nil
}
