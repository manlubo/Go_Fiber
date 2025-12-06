package repository

import (
	"context"
	"fiber/internal/entity"
	"fiber/internal/util"
	"fmt"

	driver "github.com/arangodb/go-driver"
	"github.com/google/uuid"
)

type BoardRepository struct {
	col driver.Collection
}

func NewBoardRepository(col driver.Collection) *BoardRepository {
	if col == nil {
		panic("BoardRepository: collection is nil")
	}
	return &BoardRepository{col: col}
}

// Create
func (r *BoardRepository) Create(board *entity.Board) error {
	ctx := context.Background()

	// ID 생성
	board.ID = uuid.NewString()

	_, err := r.col.CreateDocument(ctx, board)
	if err != nil {
		return err
	}

	return nil
}

// Read
func (r *BoardRepository) FindByID(id string) (*entity.Board, error) {
	ctx := context.Background()

	query := `
	FOR b IN @@col
		FILTER b.id == @id
		RETURN b
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

	var board entity.Board
	_, err = cursor.ReadDocument(ctx, &board)
	if driver.IsNotFoundGeneral(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &board, nil
}

// list
func (r *BoardRepository) List() ([]*entity.Board, error) {
	ctx := context.Background()

	query := `
	FOR b IN @@col
		RETURN b
	`

	bindVars := map[string]interface{}{
		"@col": r.col.Name(),
	}

	cursor, err := r.col.Database().Query(ctx, query, bindVars)
	if err != nil {
		return nil, err
	}
	defer cursor.Close()

	var boards []*entity.Board

	for {
		var board entity.Board
		_, err := cursor.ReadDocument(ctx, &board)

		if driver.IsNoMoreDocuments(err) {
			break
		}
		if err != nil {
			return nil, err
		}

		// 슬라이스에 추가
		boards = append(boards, &board)
	}

	return boards, nil
}

// update
func (r *BoardRepository) Update(id string, board *entity.Board) error {
	ctx := context.Background()
	fmt.Println("Update board: ", id, board.Title, board.Content)
	patch := util.BuildPatch(board, "id")

	query := `
		FOR b In @@col
			FILTER b.id == @id
			UPDATE b WITH @patch IN @@col
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
func (r *BoardRepository) Delete(id string) error {
	fmt.Println("Delete board: ", id)

	ctx := context.Background()

	query := `
		FOR b In @@col
			FILTER b.id == @id
			REMOVE { _key: b._key } IN @@col
	`

	bindVars := map[string]interface{}{
		"@col": r.col.Name(),
		"id":   id,
	}

	cursor, err := r.col.Database().Query(ctx, query, bindVars)
	if err != nil {
		return err
	}
	defer cursor.Close()

	return nil
}
