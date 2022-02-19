package db_ops

import (
	"context"

	"github.com/k0pernicus/go-photoaccess/pkg/types"
)

// PhotoExists is an easy access function to know if a Photo entity
// with a specific ID exists in DB or not
func PhotoExists(ctx context.Context, photoID string) bool {
	var photo types.Photo
	err := GetOnePhoto(ctx, photoID, &photo)
	return err == nil
}
