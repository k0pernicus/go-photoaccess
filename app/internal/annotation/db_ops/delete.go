package db_ops

import (
	"context"
	"fmt"

	"github.com/jackc/pgconn"
	app "github.com/k0pernicus/go-photoaccess/internal"
)

// DeleteOne allows to delete only one annotation based on its id (and photo_id) info
func DeleteOne(ctx context.Context, annotationID string, photoID string) (pgconn.CommandTag, error) {
	return app.DB.Exec(ctx, fmt.Sprintf("DELETE from %s where id=%s AND photo_id=%s", annotationTableName, annotationID, photoID))
}

// DeleteAllAnnotations allows to delete all annotations if they match a specific `photo_id` field
func DeleteAllAnnotations(ctx context.Context, photoID string) error {
	_, err := app.DB.Exec(ctx, fmt.Sprintf("DELETE FROM %s WHERE photo_id=%s", annotationTableName, photoID))
	if err != nil {
		return nil
	}
	return err
}
