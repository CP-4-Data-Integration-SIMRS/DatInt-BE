package usecase

import (
	"log"

	"github.com/vier21/simrs-cdc-monitoring/bin/module/log/model"
	"github.com/vier21/simrs-cdc-monitoring/bin/module/log/repository"
)

type LogUsecase interface {
	GetLogs(status, search string) ([]model.LogData, error)
}

type LogUC struct {
	repo repository.LogRepositoryInterface
}

func NewLogUsecase(repo repository.LogRepositoryInterface) *LogUC {
	return &LogUC{
		repo: repo,
	}
}

func (lu *LogUC) GetLogs(status, search string) ([]model.LogData, error) {
	logs, err := lu.repo.GetLogs(status, search)
	if err != nil {
		log.Printf("error getting logs: %s", err.Error())
		return nil, err
	}

	return logs, nil
}
