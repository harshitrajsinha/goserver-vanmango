package service

import (
	"context"

	"github.com/harshitrajsinha/goserver-vanmango/models"
)

type EngineServiceInterface interface {
	GetEngineByID(ctx context.Context, id string) (interface{}, error)
	GetAllEngine(ctx context.Context) (interface{}, error)
	CreateEngine(ctx context.Context, engineReq *models.Engine) (map[string]string, error)
	UpdateEngine(ctx context.Context, id string, engineReq *models.Engine) (int64, error)
	DeleteEngine(ctx context.Context, id string) (int64, error)
}

type VanServiceInterface interface {
	GetVanById(ctx context.Context, id string) (interface{}, error)
	GetVanByName(ctx context.Context, name string) (interface{}, error)
	GetVanByBrand(ctx context.Context, brand string) (interface{}, error)
	GetVanByCategory(ctx context.Context, category string) (interface{}, error)
	GetAllVan(ctx context.Context) (interface{}, error)
	CreateVan(ctx context.Context, vanReq *models.Van) (map[string]string, error)
	UpdateVan(ctx context.Context, vanID string, vanReq *models.Van) (int64, error)
	DeleteVan(ctx context.Context, id string) (int64, error)
}
