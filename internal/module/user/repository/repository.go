package repository

import (
	"context"
	"fiber/internal/entity"
	"fiber/internal/util"
	"fmt"
	"time"

	driver "github.com/arangodb/go-driver"
	"github.com/google/uuid"
)

type UserRepository struct {
	col driver.Collection
}

func NewUserRepository(col driver.Collection) *UserRepository {
	if col == nil {
		panic("UserRepository: collection is nil")
	}
	return &UserRepository{col: col}
}

// Create
func (r *UserRepository) Create(user *entity.User) error {
	ctx := context.Background()

	// ID 생성
	user.ID = uuid.NewString()
	user.CreatedAt = time.Now().Unix()
	user.IsActive = true

	_, err := r.col.CreateDocument(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

// Read
func (r *UserRepository) FindByID(id string) (*entity.User, error) {
	ctx := context.Background()

	query := `
	FOR u IN @@col
		FILTER u.id == @id
		FILTER u.isDeleted == false
		FILTER u.isActive == true
		RETURN u
	`

	bindVars := map[string]interface{}{
		"@col": r.col.Name(),
		"id":   id,
	}

	cursor, err := r.col.Database().Query(ctx, query, bindVars)
	if err != nil {
		return nil, err
	}
	defer cursor.Close()

	var user entity.User
	_, err = cursor.ReadDocument(ctx, &user)
	if driver.IsNotFoundGeneral(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// list
func (r *UserRepository) List() ([]*entity.User, error) {
	ctx := context.Background()

	query := `
	FOR u IN @@col
		FILTER u.isDeleted == false
		FILTER u.isActive == true
		RETURN u
	`

	bindVars := map[string]interface{}{
		"@col": r.col.Name(),
	}

	cursor, err := r.col.Database().Query(ctx, query, bindVars)
	if err != nil {
		return nil, err
	}
	defer cursor.Close()

	var users []*entity.User

	for {
		var user entity.User
		_, err := cursor.ReadDocument(ctx, &user)

		if driver.IsNoMoreDocuments(err) {
			break
		}
		if err != nil {
			return nil, err
		}

		// 슬라이스에 추가
		users = append(users, &user)
	}

	return users, nil
}

// update
func (r *UserRepository) Update(id string, user *entity.User) error {
	ctx := context.Background()
	fmt.Println("Update user: ", id, user.Name, user.Email)
	patch := util.BuildPatch(user, "id", "createdAt", "isDeleted", "deletedAt")

	query := `
		FOR u In @@col
			FILTER u.id == @id
			UPDATE u WITH @patch IN @@col
	`

	bindVars := map[string]interface{}{
		"@col":  r.col.Name(),
		"id":    id,
		"patch": patch,
	}

	cursor, err := r.col.Database().Query(ctx, query, bindVars)
	if err != nil {
		return err
	}
	defer cursor.Close()

	return nil
}

// delete
func (r *UserRepository) Delete(id string) error {
	fmt.Println("Delete user: ", id)

	ctx := context.Background()

	query := `
		FOR u In @@col
			FILTER u.id == @id
			UPDATE u WITH { 
				isActive: false,
				isDeleted: true,
				deletedAt: @deletedAt
				} IN @@col
	`

	bindVars := map[string]interface{}{
		"@col":      r.col.Name(),
		"id":        id,
		"deletedAt": time.Now().Unix(),
	}

	cursor, err := r.col.Database().Query(ctx, query, bindVars)
	if err != nil {
		return err
	}
	defer cursor.Close()

	return nil
}
