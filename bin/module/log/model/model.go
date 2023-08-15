package model

import (
	"time"

	"github.com/google/uuid"
)


type LogData struct {
    Healthcare string
    DBName     string
    TBName     string
    Status     string       
    DateTime   time.Time
    CreatedAt  time.Time
    RecordID   uuid.UUID
}


