package db_ops

import (
	"context"
	"fmt"

	app "github.com/k0pernicus/go-photoaccess/internal"
	"github.com/k0pernicus/go-photoaccess/pkg/types"
)

// CreateAnnotation allows to create one Annotation entity in the DB
func CreateAnnotation(ctx context.Context, annotation types.AnnotationCreationRequest, photoID int) (int, error) {
	insertedID := 0
	err := app.DB.QueryRow(ctx, fmt.Sprintf("INSERT INTO %s (content, photo_id, x, x2, y, y2) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", annotationTableName), annotation.Text, photoID, annotation.Coordinates.X, annotation.Coordinates.X2, annotation.Coordinates.Y, annotation.Coordinates.Y2).Scan(&insertedID)
	if err != nil {
		return insertedID, err
	}
	return insertedID, nil
}
