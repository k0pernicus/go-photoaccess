package db_ops

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	app "github.com/k0pernicus/go-photoaccess/internal"
	"github.com/k0pernicus/go-photoaccess/pkg/types"
	log "github.com/sirupsen/logrus"
)

func CountRows(ctx context.Context) (int, error) {
	var count int // Default to 0
	if err := app.DB.QueryRow(ctx, fmt.Sprintf("SELECT count(*) FROM %s", photosTableName)).Scan(&count); err != nil && err != pgx.ErrNoRows {
		return 0, err
	}
	return count, nil
}

// GetOnePhoto returns the result of the 'GET' operation with id checking
func GetOnePhoto(ctx context.Context, id string, photo *types.Photo) error {
	return app.DB.QueryRow(ctx, fmt.Sprintf("SELECT * FROM %s WHERE id = %s", photosTableName, id)).Scan(&photo.ID, &photo.Content, &photo.CreatedAt, &photo.UpdatedAt)
}

// GetAllPhotos returns the result of a 'GET' operation for possible multiple elements
// and for Photo type only
func GetAllPhotos(ctx context.Context, photos *[]types.Photo) error {
	rows, err := app.DB.Query(ctx, fmt.Sprintf("SELECT id, content, created_at, updated_at FROM %s", photosTableName))
	if err != nil {
		return err
	}
	defer rows.Close()

	i := 0
	for rows.Next() {
		var p types.Photo
		err := rows.Scan(&p.ID, &p.Content, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			log.Warningf("cannot scan annotation with id due to error: %+v\n", err)
			continue
		}
		(*photos)[i] = p
		i++
	}

	return nil
}

// GetAllPhotos returns the result of a 'GET' operation for possible multiple elements
// and for Photo type only
func GetAllPhotosWithID(ctx context.Context, key string, value string) ([]types.Photo, error) {
	rows, err := app.DB.Query(ctx, fmt.Sprintf("SELECT * FROM %s WHERE %s = %s", photosTableName, key, value))
	photos := []types.Photo{}
	if err != nil {
		return photos, err
	}
	defer rows.Close()

	for rows.Next() {
		var p types.Photo
		err := rows.Scan(&p.ID, &p.Content, &p.CreatedAt)
		if err != nil {
			log.Warningf("cannot scan annotation with id due to error: %+v\n", err)
			continue
		}
		photos = append(photos, p)
	}

	return photos, nil
}
