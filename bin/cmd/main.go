package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
	lh "github.com/vier21/simrs-cdc-monitoring/bin/module/log/handler"
	lm "github.com/vier21/simrs-cdc-monitoring/bin/module/monitor/handler"
	"github.com/vier21/simrs-cdc-monitoring/bin/pkg/elastic"
	"github.com/vier21/simrs-cdc-monitoring/bin/pkg/mysql"
	"github.com/vier21/simrs-cdc-monitoring/config"
)

func main() {
	mysql.InitDB()
	elastic.InitElastic()

	m := chi.NewRouter()

	RunServer(m)

	server := http.Server{
		Addr:         config.GetConfig().ServerPort,
		Handler:      m,
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		if err := server.ListenAndServe(); err == http.ErrServerClosed {
			log.Fatalf("error starting server: ",err.Error())
			return
		}
	}()
}

func RunServer(c *chi.Mux) {
	lh.InitLogHttpHandler(c)
	lm.InitMonitorHttpHandler(c)

}
