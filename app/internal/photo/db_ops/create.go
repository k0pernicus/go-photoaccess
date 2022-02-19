package db_ops

import (
	"context"
	"fmt"

	app "github.com/k0pernicus/go-photoaccess/internal"
	"github.com/k0pernicus/go-photoaccess/pkg/types"
)

func CreatePhoto(ctx context.Context, photo types.PhotoCreationRequest) (int, error) {
	if photo.IsAdditional {
		return createAdditionalPhoto(ctx, photo)
	}
	return createPhoto(ctx, photo)
}

func createPhoto(ctx context.Context, photo types.PhotoCreationRequest) (int, error) {
	insertedID := 0
	err := app.DB.QueryRow(ctx, fmt.Sprintf("INSERT INTO %s (content, is_additional, annotation_id) VALUES ($1, false, 0) RETURNING id", photosTableName), photo.Data).Scan(&insertedID)
	if err != nil {
		return insertedID, err
	}
	return insertedID, nil
}

func createAdditionalPhoto(ctx context.Context, photo types.PhotoCreationRequest) (int, error) {
	insertedID := 0
	err := app.DB.QueryRow(ctx, fmt.Sprintf("INSERT INTO %s (content, is_additional, annotation_id) VALUES ($1, $2, $3) RETURNING id", photosTableName), photo.Data, true, photo.AnnotationID).Scan(&insertedID)
	if err != nil {
		return insertedID, err
	}
	return insertedID, nil
}
