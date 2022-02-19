package db_ops

import (
	"context"
	"fmt"

	app "github.com/k0pernicus/go-photoaccess/internal"
	"github.com/k0pernicus/go-photoaccess/pkg/types"
)

// CreatePhoto allows to create one Photo entity in the DB
func CreatePhoto(ctx context.Context, photo types.PhotoCreationRequest) (int, error) {
	insertedID := 0
	err := app.DB.QueryRow(ctx, fmt.Sprintf("INSERT INTO %s (content) VALUES ($1) RETURNING id", photosTableName), photo.Data).Scan(&insertedID)
	if err != nil {
		return insertedID, err
	}
	return insertedID, nil
}
