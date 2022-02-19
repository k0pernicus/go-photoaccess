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

// GetOneAnnotation returns the result of the 'GET' operation with id checking
func GetOneAnnotation(ctx context.Context, id string, annotation *types.Annotation) error {
	return app.DB.QueryRow(ctx, fmt.Sprintf("SELECT * FROM %s WHERE id = %s", annotationTableName, id)).Scan(annotation.ID, annotation.Content, annotation.PhotoID, annotation.CreatedAt, annotation.UpdatedAt)
}

// GetAllAnnotations returns the result of a 'GET' operation for possible multiple elements
// and for Annotation type only
func GetAllAnnotations(ctx context.Context, annotations *[]types.Annotation, photoID string) error {
	rows, err := app.DB.Query(ctx, fmt.Sprintf("SELECT * FROM %s WHERE photo_id=%s", annotationTableName, photoID))
	if err != nil {
		return err
	}
	defer rows.Close()

	i := 0
	for rows.Next() {
		var a types.Annotation
		err := rows.Scan(&a.ID, &a.Content, &a.CreatedAt)
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
func GetAllAnnotationsWithID(ctx context.Context, key string, value string) ([]types.Annotation, error) {
	rows, err := app.DB.Query(ctx, fmt.Sprintf("SELECT * FROM %s WHERE %s = %s", annotationTableName, key, value))
	annotations := []types.Annotation{}
	if err != nil {
		return annotations, err
	}
	defer rows.Close()

	for rows.Next() {
		var a types.Annotation
		err := rows.Scan(&a.ID, &a.Content, &a.CreatedAt)
		if err != nil {
			log.Warningf("cannot scan annotation with id due to error: %+v\n", err)
			continue
		}
		annotations = append(annotations, a)
	}

	return annotations, nil
}
