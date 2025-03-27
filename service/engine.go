package service

import (
	"context"

	"github.com/harshitrajsinha/van-man-go/models"
	"github.com/harshitrajsinha/van-man-go/store"
)

type EngineService struct {
	store store.EngineStoreInterface
}

func NewEngineService(store store.EngineStoreInterface) *EngineService{
	return &EngineService{
		store: store,
	}
}

func (s *EngineService) GetEngineByID(ctx context.Context, id string) (interface{}, error){
	engine, err := s.store.EngineById(ctx, id)
	if err != nil{
		return nil, err
	}
	return &engine, nil
}

func (s *EngineService) CreateEngine(ctx context.Context, engineReq *models.Engine)(interface{}, error){
	if err := models.ValidateEngineReq(*engineReq); err != nil{
		return nil, err
	}

	createdEngine, err := s.store.CreateEngine(ctx, engineReq)
	if err != nil{
		return nil, err
	}

	return &createdEngine, nil
}

func (s *EngineService) UpdateEngine(ctx context.Context, id string, engineReq *models.Engine)(interface{}, error){
	// if err := models.ValidateEngineReq(*engineReq); err != nil{
	// 	return nil, err
	// }

	updatedEngine, err := s.store.EngineUpdate(ctx, id, engineReq)
	if err != nil{
		return nil, err
	}
	return &updatedEngine, nil
}


func (s *EngineService) DeleteEngine(ctx context.Context, id string)(interface{}, error){

	deletedEngine, err := s.store.EngineDelete(ctx, id)
	if err != nil{
		return nil, err
	}
	return &deletedEngine, nil
}