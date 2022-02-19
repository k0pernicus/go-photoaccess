package db_ops

import (
	"context"
	"strconv"

	"github.com/k0pernicus/go-photoaccess/pkg/types"
)

func AnnotationExists(ctx context.Context, annotationID int) bool {
	var annotation types.Annotation
	err := GetOneAnnotation(ctx, strconv.Itoa(annotationID), &annotation)
	return err == nil
}
