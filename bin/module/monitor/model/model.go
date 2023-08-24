package model

type DatabaseInfo struct {
	ID string `json:"id,omitempty"`
	DBName     string `json:"dbname"`
	TotalTable int `json:"totalTable"`
	TableInfo  []Table `json:"tableInfo"`
}

type Table struct {
	TableName       string `json:"tableName"`
	TotalRecord     int `json:"totalRecord"`
	NewData         int `json:"newData"`
	DeltaData       int `json:"deltaData"`
	CurrentCaptured int `json:"currentCapured"`
}
