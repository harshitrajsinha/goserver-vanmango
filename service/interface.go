package service

import (
	"context"

	"github.com/harshitrajsinha/van-man-go/models"
)

type EngineServiceInterface interface {
	GetEngineByID(ctx context.Context, id string) (interface{}, error)
	CreateEngine(ctx context.Context, engineReq *models.Engine) (interface{}, error)
	UpdateEngine(ctx context.Context, id string, engineReq *models.Engine) (interface{}, error)
	DeleteEngine(ctx context.Context, id string) (interface{}, error)
}