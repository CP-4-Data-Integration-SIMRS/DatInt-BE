package kafka

import (
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/vier21/simrs-cdc-monitoring/bin/module/monitor/model"
)

func (p *Producer) RunCron() {
	s := gocron.NewScheduler(time.Now().Local().Location())
	s.Every(3).Second().Do(p.collectData)
	s.StartAsync()
}

func (p *Producer) collectData() {
	err := p.SendDataToProducer([]model.DatabaseInfo{
		{
			ID: "1",
			DBName:     "rs_advent",
			TotalTable: 123,
			TableInfo: []model.Table{
				{
					TableName:       "pasien",
					TotalRecord:     20,
					NewData:         2,
					DeltaData:       3,
					CurrentCaptured: 12,
				},
				{
					TableName:       "admin",
					TotalRecord:     212,
					NewData:         2,
					DeltaData:       3,
					CurrentCaptured: 12,
				},
				{
					TableName:       "poli",
					TotalRecord:     20,
					NewData:         2,
					DeltaData:       3,
					CurrentCaptured: 12,
				},
			},
		},
		{
			ID: "2",
			DBName:     "rs_hermina",
			TotalTable: 123,
			TableInfo: []model.Table{
				{
					TableName:       "pasien",
					TotalRecord:     20,
					NewData:         2,
					DeltaData:       3,
					CurrentCaptured: 12,
				},
				{
					TableName:       "admin",
					TotalRecord:     20,
					NewData:         2,
					DeltaData:       3,
					CurrentCaptured: 12,
				},
				{
					TableName:       "poli",
					TotalRecord:     20,
					NewData:         2,
					DeltaData:       3,
					CurrentCaptured: 12,
				},
			},
		},
		{
			ID: "3",
			DBName:     "rsud",
			TotalTable: 123,
			TableInfo: []model.Table{
				{
					TableName:       "pasien",
					TotalRecord:     20,
					NewData:         2,
					DeltaData:       3,
					CurrentCaptured: 12,
				},
				{
					TableName:       "admin",
					TotalRecord:     20,
					NewData:         2,
					DeltaData:       3,
					CurrentCaptured: 12,
				},
				{
					TableName:       "poli",
					TotalRecord:     20,
					NewData:         2,
					DeltaData:       3,
					CurrentCaptured: 12,
				},
			},
		},
		
	})

	if err != nil {
		log.Println(err)
		return
	}
}
