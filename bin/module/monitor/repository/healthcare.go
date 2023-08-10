package repository

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/vier21/simrs-cdc-monitoring/bin/pkg/mysql"
)

type HCRepositoryInterface interface {
	GetAllDB() ([]string, error)
	GetAllTableByDB(dbname string) ([]string, error)
	CountTotalTablesByDBName(dbname string) (int, error)
	CountTableRecords(dbname string, tablename string) (int, error)
	CountNewData(dbname string) (int, error)
	CountDeltaData(dbname string) (int, error)
}

type HCRepository struct {
	db *sqlx.DB
}

func NewHealthCareRepository() *HCRepository {
	return &HCRepository{
		db: mysql.DB,
	}
}

func (h *HCRepository) GetAllDB() ([]string, error) {
	var dbs []string

	if err := h.db.Select(&dbs, "SHOW DATABASES"); err != nil {
		log.Println(err)
		return []string{}, fmt.Errorf("error get all db %s", err.Error())
	}

	return dbs, nil
}

func (h *HCRepository) CountTotalTablesByDBName(dbname string) (int, error) {
	var count int

	if err := h.db.Get(&count, "SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = ?", dbname); err != nil {
		log.Println(err)
		return 0, fmt.Errorf("error count table %s", err.Error())
	}

	return count, nil
}

func (h *HCRepository) CountTableRecords(dbname string, tablename string) (int, error) {
	totalrecord := 0
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s.%s", dbname, tablename)

	if err := h.db.Get(&totalrecord, query); err != nil {
		log.Println(err)

		return 0, err
	}

	return totalrecord, nil
}

func (h *HCRepository) GetAllTableByDB(dbname string) ([]string, error) {
	var tbnames []string

	if err := h.db.Select(&tbnames, "SELECT TABLE_NAME FROM information_schema.TABLES WHERE TABLE_SCHEMA = ?", dbname); err != nil {
		log.Println(err)
		return []string{}, err
	}

	return tbnames, nil
}

func (h *HCRepository) CountNewData(dbname string) (int, error) {

	return 0, nil
}

func (h *HCRepository) CountDeltaData(dbname string) (int, error) {
	return 0, nil
}
