package usecase

import (
	"github.com/vier21/simrs-cdc-monitoring/bin/module/monitor/model"
	"github.com/vier21/simrs-cdc-monitoring/bin/module/monitor/repository"
)

type MonitoringUsecase interface {
	GetAllDatabaseInfo() ([]model.DatabaseInfo, error)
	GetDBTableInfo(dbname string) ([]model.Table, error)
}

type HCUsecase struct {
	repo repository.HCRepositoryInterface
}

func NewMonitorUsecase(repo repository.HCRepositoryInterface) *HCUsecase {
	return &HCUsecase{
		repo: repo,
	}
}

func (hu *HCUsecase) GetAllDatabaseInfo() ([]model.DatabaseInfo, error) {
	return []model.DatabaseInfo{}, nil

}

func (hu *HCUsecase) GetDBTableInfo(dbname string) ([]model.Table, error) {
	return []model.Table{}, nil
}
