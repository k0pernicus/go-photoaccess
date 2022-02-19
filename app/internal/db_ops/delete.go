package db_ops

import (
	"context"
	"fmt"

	app "github.com/k0pernicus/go-photoaccess/internal"
)

// TODO
func DeleteAllAnnotations(ctx context.Context, photoID string) error {
	_, err := app.DB.Exec(ctx, fmt.Sprintf("DELETE FROM ANNOTATIONS WHERE photo_id=%s", photoID))
	if err != nil {
		return nil
	}
	// TODO: Delete all additional photos
	return err
}

func DeleteAdditionalPhotos(ctx context.Context, annotationID string) error {
	_, err := app.DB.Exec(ctx, fmt.Sprintf("DELETE FROM PHOTOS WHERE is_additional=true AND annotation_id=%s", annotationID))
	return err
}
