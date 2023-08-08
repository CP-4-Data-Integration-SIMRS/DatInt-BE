package model

import (
	"time"

	"github.com/google/uuid"
)

type LogData struct {
	Healthcare string
	DBName     string
	CreatedAt  time.Time
	RecordId   uuid.UUID
}



