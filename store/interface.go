package store

import (
	"context"

	"github.com/harshitrajsinha/van-man-go/models"
)

type EngineStoreInterface interface {
	EngineById(ctx context.Context, id string) (interface{}, error)
	CreateEngine(ctx context.Context, engineReq *models.Engine) (interface{}, error)
	EngineUpdate(ctx context.Context, id string, engineReq *models.Engine) (interface{}, error)
	EngineDelete(ctx context.Context, id string) (interface{}, error)
}