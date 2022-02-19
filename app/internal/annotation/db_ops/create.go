package db_ops

import (
	"context"
	"fmt"

	app "github.com/k0pernicus/go-photoaccess/internal"
	"github.com/k0pernicus/go-photoaccess/pkg/types"
)

func CreateAnnotation(ctx context.Context, annotation types.AnnotationCreationRequest, photoID int) (int, error) {
	insertedID := 0
	err := app.DB.QueryRow(ctx, fmt.Sprintf("INSERT INTO %s (content, photo_id) VALUES ($1, $2) RETURNING id", annotationTableName), annotation.Text, photoID).Scan(&insertedID)
	if err != nil {
		return insertedID, err
	}
	return insertedID, nil
}
