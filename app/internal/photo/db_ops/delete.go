package db_ops

import (
	"context"
	"fmt"

	"github.com/jackc/pgconn"
	app "github.com/k0pernicus/go-photoaccess/internal"
)

func DeleteOne(ctx context.Context, photoID string) (pgconn.CommandTag, error) {
	return app.DB.Exec(ctx, fmt.Sprintf("DELETE from %s where id=%s", photosTableName, photoID))
}
