package uploadprovider

import (
	"context"
	"go-simple-service/common"
)

type UploadProvider interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error)
}
