package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
	hl "github.com/vier21/simrs-cdc-monitoring/bin/module/log/handler"
	repoLog "github.com/vier21/simrs-cdc-monitoring/bin/module/log/repository"
	usecaseLog "github.com/vier21/simrs-cdc-monitoring/bin/module/log/usecase"
	hm "github.com/vier21/simrs-cdc-monitoring/bin/module/monitor/handler"
	"github.com/vier21/simrs-cdc-monitoring/bin/module/monitor/repository"
	"github.com/vier21/simrs-cdc-monitoring/bin/module/monitor/usecase"
	"github.com/vier21/simrs-cdc-monitoring/bin/pkg/elastic"
)

func main() {
	elastic.InitElastic()

	m := chi.NewRouter()

	RunServer(m)

	port := envPortOr("3030")

	server := http.Server{
		Addr:         port,
		Handler:      m,
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Server start on %s \n", server.Addr)

	if err := server.ListenAndServe(); err == http.ErrServerClosed {
		log.Fatalf("error starting server: %s", err.Error())
		return
	}
}

func RunServer(c *chi.Mux) {

	monitorRepo := repository.NewHealthCareRepository()
	monitorUsecase := usecase.NewMonitorUsecase(monitorRepo)
	hm.InitMonitorHttpHandler(c, monitorUsecase)

	logRepo := repoLog.NewLogRepository()
	logUsecase := usecaseLog.NewLogUsecase(logRepo)
	hl.InitLogHttpHandler(c, logUsecase)

}

func envPortOr(port string) string {
	// If `PORT` variable in environment exists, return it
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	// Otherwise, return the value of `port` variable from function argument
	return ":" + port
}
