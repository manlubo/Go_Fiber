package repository

import (
	"context"
	"fiber/internal/entity"
	"fiber/internal/util"
	"time"

	driver "github.com/arangodb/go-driver"
	"github.com/google/uuid"
)

type ReplyRepository struct {
	col driver.Collection
}

func NewReplyRepository(col driver.Collection) *ReplyRepository {
	if col == nil {
		panic("ReplyRepository: collection is nil")
	}
	return &ReplyRepository{col: col}
}

// Create
func (r *ReplyRepository) Create(reply *entity.Reply) error {
	ctx := context.Background()

	reply.ID = uuid.NewString()
	reply.CreatedAt = time.Now().Unix()
	reply.IsActive = true

	_, err := r.col.CreateDocument(ctx, reply)
	if err != nil {
		return err
	}

	return nil
}

// Read
func (r *ReplyRepository) FindByID(id string) (*entity.Reply, error) {
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

	var reply entity.Reply
	_, err = cursor.ReadDocument(ctx, &reply)
	if driver.IsNotFoundGeneral(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &reply, nil
}

// list

func (r *ReplyRepository) List() ([]*entity.Reply, error) {
	ctx := context.Background()

	query := `
		FOR r IN @@col
			FILTER r.isDeleted == false
			FILTER r.isActive == true
			RETURN r
	`

	bindVars := map[string]interface{}{
		"@col": r.col.Name(),
	}

	cursor, err := r.col.Database().Query(ctx, query, bindVars)
	if err != nil {
		return nil, err
	}
	defer cursor.Close()

	var replys []*entity.Reply

	for {
		var reply entity.Reply
		_, err := cursor.ReadDocument(ctx, &reply)

		if driver.IsNoMoreDocuments(err) {
			break
		}
		if err != nil {
			return nil, err
		}

		replys = append(replys, &reply)
	}

	return replys, nil
}

// update
func (r *ReplyRepository) Update(id string, reply *entity.Reply) error {
	ctx := context.Background()
	patch := util.BuildPatch(reply, "id", "createdAt", "isDeleted", "deletedAt")

	query := `
		FOR u IN @@col
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
func (r *ReplyRepository) Delete(id string) error {
	ctx := context.Background()

	query := `
		FOR r IN @@col
			FILTER r.id == @id
			UPDATE r WITH { 
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
