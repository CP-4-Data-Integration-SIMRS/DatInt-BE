package usecase

import (
	"log"

	"github.com/vier21/simrs-cdc-monitoring/bin/module/log/model"
	"github.com/vier21/simrs-cdc-monitoring/bin/module/log/repository"
)

type LogUsecase interface {
	GetLogs() ([]model.LogData, error)
}

type LogUC struct {
	repo repository.LogRepositoryInterface
}

func NewLogUsecase(repo repository.LogRepositoryInterface) *LogUC {
	return &LogUC{
		repo: repo,
	}
}

func (lu *LogUC) GetLogs() ([]model.LogData, error) {
	logs, err := lu.repo.GetLogs()
	if err != nil {
		log.Printf("error getting logs: %s", err.Error())
		return nil, err
	}
	return logs, nil
}
