package main

import (
	"net/http"
	"os"

	"github.com/Lanquill/Forge/pkg/auth"
	"github.com/Lanquill/Forge/pkg/queue"
	"github.com/Lanquill/Forge/router"
	"github.com/Lanquill/go-logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

var (
	STATICPATH = os.Getenv("LQ_STATIC_PATH")
)

func main() {

	logPath := os.Getenv("LQ_DB_PATH") + "/logs/forge/forge.log"
	logLevel := os.Getenv("LOG_LEVEL")
	log := logger.Get(logPath, logLevel)

	// Open rabbitMQ connection
	queue.InitializeQueue()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(auth.Auth)
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir(STATICPATH))))

	router.InitRoutes(r)
	log.Info("Forge Running!")
	err := http.ListenAndServe(":2005", r)
	if err != nil {
		log.Fatal("Error initializing routes: ", zap.Error(err))
	}
}
