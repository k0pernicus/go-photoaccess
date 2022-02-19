package db_ops

import (
	"context"
	"fmt"

	"github.com/jackc/pgconn"
	app "github.com/k0pernicus/go-photoaccess/internal"
)

func DeleteOne(ctx context.Context, annotationID string) (pgconn.CommandTag, error) {
	return app.DB.Exec(ctx, fmt.Sprintf("DELETE from %s where id=%s", annotationTableName, annotationID))
}

// TODO
func DeleteAllAnnotations(ctx context.Context, photoID string) error {
	_, err := app.DB.Exec(ctx, fmt.Sprintf("DELETE FROM %s WHERE photo_id=%s", annotationTableName, photoID))
	if err != nil {
		return nil
	}
	// TODO: Delete all additional photos
	return err
}
