package image

import (
	"context"
)

type Storage interface {
	Save(ctx context.Context, filename string, data []byte) (string, error)
}
