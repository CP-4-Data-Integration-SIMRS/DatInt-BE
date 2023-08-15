package model

import (
	"time"

	"github.com/google/uuid"
)


type LogData struct {
    Healthcare string
    DBName     string
    TBName     string
    Status     string       // Saya menambahkan field Status di sini
    DateTime   time.Time
    CreatedAt  time.Time
    RecordId   uuid.UUID
}


