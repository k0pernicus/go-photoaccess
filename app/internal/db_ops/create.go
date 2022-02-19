package db_ops

import (
	"context"
	"fmt"

	app "github.com/k0pernicus/go-photoaccess/internal"
	"github.com/k0pernicus/go-photoaccess/pkg/types"
)

func CreatePhoto(ctx context.Context, tableName string, photo types.PhotoCreationRequest) (int, error) {
	if photo.IsAdditional {
		return createAdditionalPhoto(ctx, tableName, photo)
	}
	return createPhoto(ctx, tableName, photo)
}

func createPhoto(ctx context.Context, tableName string, photo types.PhotoCreationRequest) (int, error) {
	insertedID := 0
	err := app.DB.QueryRow(ctx, fmt.Sprintf("INSERT INTO %s (content, is_additional, annotation_id) VALUES ($1, false, 0) RETURNING id", tableName), photo.Data).Scan(&insertedID)
	if err != nil {
		return insertedID, err
	}
	return insertedID, nil
}

func createAdditionalPhoto(ctx context.Context, tableName string, photo types.PhotoCreationRequest) (int, error) {
	insertedID := 0
	err := app.DB.QueryRow(ctx, fmt.Sprintf("INSERT INTO %s (content, is_additional, annotation_id) VALUES ($1, $2, $3) RETURNING id", tableName), photo.Data, true, photo.AnnotationID).Scan(&insertedID)
	if err != nil {
		return insertedID, err
	}
	return insertedID, nil
}

func CreateAnnotation(ctx context.Context, tableName string, annotation types.AnnotationCreationRequest, photoID int) (int, error) {
	insertedID := 0
	err := app.DB.QueryRow(ctx, fmt.Sprintf("INSERT INTO %s (content, photo_id) VALUES ($1, $2) RETURNING id", tableName), annotation.Text, photoID).Scan(&insertedID)
	if err != nil {
		return insertedID, err
	}
	return insertedID, nil
}
