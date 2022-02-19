package db_ops

import (
	"context"
	"fmt"

	app "github.com/k0pernicus/go-photoaccess/internal"
	"github.com/k0pernicus/go-photoaccess/pkg/types"
	log "github.com/sirupsen/logrus"
)

// GetOnePhoto returns the result of the 'GET' operation with id checking
func GetOnePhoto(ctx context.Context, tableName string, id string, photo *types.Photo) error {
	return app.DB.QueryRow(ctx, fmt.Sprintf("SELECT * FROM %s WHERE ID = %s", tableName, id)).Scan(&photo.ID, &photo.Content, &photo.AnnotationID, &photo.IsAdditional, &photo.CreatedAt, &photo.UpdatedAt)
}

// GetOneAnnotation returns the result of the 'GET' operation with id checking
func GetOneAnnotation(ctx context.Context, tableName string, id string, annotation *types.Annotation) error {
	return app.DB.QueryRow(ctx, fmt.Sprintf("SELECT * FROM %s WHERE ID = %s", tableName, id)).Scan(annotation.ID, annotation.Content, annotation.PhotoID, annotation.CreatedAt, annotation.UpdatedAt)
}

// GetAllAnnotations returns the result of a 'GET' operation for possible multiple elements
// and for Annotation type only
func GetAllAnnotations(ctx context.Context, tableName string, key string, value string) ([]types.Annotation, error) {
	rows, err := app.DB.Query(ctx, fmt.Sprintf("SELECT * FROM %s WHERE %s = %s", tableName, key, value))
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

// GetAllPhotos returns the result of a 'GET' operation for possible multiple elements
// and for Photo type only
func GetAllPhotos(ctx context.Context, tableName string, key string, value string) ([]types.Photo, error) {
	rows, err := app.DB.Query(ctx, fmt.Sprintf("SELECT * FROM %s WHERE %s = %s", tableName, key, value))
	photos := []types.Photo{}
	if err != nil {
		return photos, err
	}
	defer rows.Close()

	for rows.Next() {
		var p types.Photo
		err := rows.Scan(&p.ID, &p.Content, &p.IsAdditional, &p.CreatedAt)
		if err != nil {
			log.Warningf("cannot scan annotation with id due to error: %+v\n", err)
			continue
		}
		photos = append(photos, p)
	}

	return photos, nil
}
