package db_ops

import (
	"context"

	"github.com/k0pernicus/go-photoaccess/pkg/types"
)

func PhotoExists(ctx context.Context, photoID string) bool {
	var photo types.Photo
	err := GetOnePhoto(ctx, photoID, &photo)
	return err == nil
}
