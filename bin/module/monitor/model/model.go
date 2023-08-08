package model

type DatabaseInfo struct {
	DBName          string
	TableInfo       []Table
	TotalTable      int
	NewData         int
	DeltaData       int
	CurrentCaptured int
}

type Table struct {
	TableName   string
	TotalRecord int
}
