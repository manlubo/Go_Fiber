//go:build wireinject
// +build wireinject

package user

import (
	"fiber/internal/db"
	"fiber/internal/module/user/controller"
	"fiber/internal/module/user/repository"
	"fiber/internal/module/user/service"

	driver "github.com/arangodb/go-driver"
	"github.com/google/wire"
)

// User 컬렉션 Provider
func provideUserCollection(d driver.Database) driver.Collection {
	return db.EnsureCollection(d, "users")
}

// Wire injector
func InitializeUserModule() *UserModule {
	wire.Build(
		db.ConnectArango,
		provideUserCollection,
		repository.NewUserRepository,
		service.NewUserService,
		controller.NewUserController,
		NewUserModule,
	)
	return nil
}
