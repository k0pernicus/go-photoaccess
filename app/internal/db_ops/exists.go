package db_ops

import (
	"context"
	"strconv"

	"github.com/k0pernicus/go-photoaccess/pkg/types"
)

func AnnotationExists(ctx context.Context, annotationID int) bool {
	var annotation types.Annotation
	err := GetOneAnnotation(ctx, "ANNOTATIONS", strconv.Itoa(annotationID), &annotation)
	return err == nil
}

func PhotoExists(ctx context.Context, photoID string) bool {
	var photo types.Photo
	err := GetOnePhoto(ctx, "PHOTOS", photoID, &photo)
	return err == nil
}
