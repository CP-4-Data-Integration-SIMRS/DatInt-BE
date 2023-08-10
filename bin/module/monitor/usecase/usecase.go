package usecase

import (
	"log"

	"github.com/vier21/simrs-cdc-monitoring/bin/module/monitor/model"
	"github.com/vier21/simrs-cdc-monitoring/bin/module/monitor/repository"
)

type MonitoringUsecase interface {
	GetAllDatabaseInfo() ([]model.DatabaseInfo, error)
	GetDBTableInfo(dbname, tbname string) ([]model.Table, error)
	GetAllTableInfoByDB(dbname string) ([]model.Table, error)
	GetDBInfo(dbname string) (model.DatabaseInfo, error)
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
	databases, err := hu.repo.GetAllDB()

	dbsInfo := []model.DatabaseInfo{}

	if err != nil {
		log.Println(err)
		return []model.DatabaseInfo{}, err
	}

	for i := range databases {
		if databases[i] != "information_schema" && databases[i] != "performance_schema" && databases[i] != "phpmyadmin" && databases[i] != "mysql" {
			dbInfo, err := hu.GetDBInfo(databases[i])
			if err != nil {
				log.Printf("error gettting all db info: %s", err.Error())
				return []model.DatabaseInfo{}, err
			}

			if dbInfo.DBName != "" {

				dbsInfo = append(dbsInfo, dbInfo)
			}
		}
	}
	return dbsInfo, nil
}

func (hu *HCUsecase) GetDBTableInfo(dbname, tbname string) (model.Table, error) {
	recordsCount, err := hu.repo.CountTableRecords(dbname, tbname)

	if err != nil {
		log.Printf("error get table info: %s", err.Error())
		return model.Table{}, err
	}

	tableinfo := model.Table{
		TableName:       tbname,
		TotalRecord:     recordsCount,
		NewData:         40,
		DeltaData:       20,
		CurrentCaptured: 200,
	}

	return tableinfo, nil
}

func (hu *HCUsecase) GetAllTableInfoByDB(dbname string) ([]model.Table, error) {
	var tablesInfo []model.Table
	tables, err := hu.repo.GetAllTableByDB(dbname)

	if err != nil {
		log.Printf("error get table info: %s", err.Error())
		return []model.Table{}, err
	}

	for i := range tables {
		table, err := hu.GetDBTableInfo(dbname, tables[i])
		if err != nil {
			log.Printf("error get table info: %s", err.Error())
			return []model.Table{}, err
		}

		tablesInfo = append(tablesInfo, table)

	}
	return tablesInfo, nil
}

func (hu *HCUsecase) GetDBInfo(dbname string) (model.DatabaseInfo, error) {
	tableinfo, err := hu.GetAllTableInfoByDB(dbname)
	if err != nil {
		log.Printf("error get dbinfo info: %s", err.Error())
		return model.DatabaseInfo{}, err
	}

	totalTable, err := hu.repo.CountTotalTablesByDBName(dbname)
	if err != nil {
		log.Printf("error get dbinfo info: %s", err.Error())
		return model.DatabaseInfo{}, err
	}

	return model.DatabaseInfo{
		DBName:     dbname,
		TotalTable: totalTable,
		TableInfo:  tableinfo,
	}, nil
}
