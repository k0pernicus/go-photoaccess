package db_ops

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	app "github.com/k0pernicus/go-photoaccess/internal"
	"github.com/k0pernicus/go-photoaccess/pkg/types"
	log "github.com/sirupsen/logrus"
)

func CountRows(ctx context.Context, photoID string) (int, error) {
	var count int // Default to 0
	if err := app.DB.QueryRow(ctx, fmt.Sprintf("SELECT count(*) FROM %s WHERE photo_id=%s", annotationTableName, photoID)).Scan(&count); err != nil && err != pgx.ErrNoRows {
		return 0, err
	}
	return count, nil
}

// GetOneAnnotation retrieves the Annotation entity that asserts the id passed as parameter
func GetOneAnnotation(ctx context.Context, id string, annotation *types.Annotation) error {
	return app.DB.QueryRow(ctx, fmt.Sprintf("SELECT id, content, x, x2, y, y2, photo_id, created_at, updated_at FROM %s WHERE id = %s", annotationTableName, id)).Scan(&annotation.ID, &annotation.Content, &annotation.Coordinates.X, &annotation.Coordinates.X2, &annotation.Coordinates.Y, &annotation.Coordinates.Y2, &annotation.PhotoID, &annotation.CreatedAt, &annotation.UpdatedAt)
}

// GetOneAnnotationWithKnownPhoto returns the result of the 'GET' operation with id checking of the annotation AND associated photo
func GetOneAnnotationWithKnownPhoto(ctx context.Context, annotationID string, photoID string, annotation *types.Annotation) error {
	return app.DB.QueryRow(ctx, fmt.Sprintf("SELECT id, content, x, x2, y, y2, photo_id, created_at, updated_at FROM %s WHERE id = %s AND photo_id=%s", annotationTableName, annotationID, photoID)).Scan(&annotation.ID, &annotation.Content, &annotation.Coordinates.X, &annotation.Coordinates.X2, &annotation.Coordinates.Y, &annotation.Coordinates.Y2, &annotation.PhotoID, &annotation.CreatedAt, &annotation.UpdatedAt)
}

// GetAllAnnotations returns the result of a 'GET' operation for possible multiple elements
// and for Annotation type only
func GetAllAnnotations(ctx context.Context, annotations *[]types.Annotation, photoID string) error {
	rows, err := app.DB.Query(ctx, fmt.Sprintf("SELECT id, content, x, x2, y, y2, photo_id, created_at, updated_at FROM %s WHERE photo_id=%s", annotationTableName, photoID))
	if err != nil {
		return err
	}
	defer rows.Close()

	i := 0
	for rows.Next() {
		var a types.Annotation
		err := rows.Scan(&a.ID, &a.Content, &a.Coordinates.X, &a.Coordinates.X2, &a.Coordinates.Y, &a.Coordinates.Y2, &a.PhotoID, &a.CreatedAt, &a.UpdatedAt)
		if err != nil {
			log.Warningf("cannot scan annotation with id due to error: %+v\n", err)
			continue
		}
		(*annotations)[i] = a
	}

	return nil
}

// GetAllAnnotations returns the result of a 'GET' operation for possible multiple elements
// and for Annotation type only
func GetAllAnnotationsWithPhotoID(ctx context.Context, photoID string) ([]types.Annotation, error) {
	rows, err := app.DB.Query(ctx, fmt.Sprintf("SELECT id, content, x, x2, y, y2, created_at, updated_at FROM %s WHERE photo_id = %s", annotationTableName, photoID))
	annotations := []types.Annotation{}
	if err != nil {
		return annotations, err
	}
	defer rows.Close()

	for rows.Next() {
		var a types.Annotation
		err := rows.Scan(&a.ID, &a.Content, &a.Coordinates.X, &a.Coordinates.X2, &a.Coordinates.Y, &a.Coordinates.Y2, &a.CreatedAt, &a.UpdatedAt)
		if err != nil {
			log.Warningf("cannot scan annotation with id due to error: %+v\n", err)
			continue
		}
		annotations = append(annotations, a)
	}

	return annotations, nil
}
