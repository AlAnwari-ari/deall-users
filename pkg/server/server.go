package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	cfg "github.com/deall-users/config"
	"github.com/deall-users/internal/module"
	r "github.com/deall-users/pkg/gin"
	"github.com/deall-users/pkg/gorm"
	"github.com/deall-users/pkg/migration"
	"github.com/gin-gonic/gin"
)

// create new server and initialize machine or other external infra
func NewServer() (*Server, error) {
	db, err := gorm.NewPostgresClient()
	if err != nil {
		return nil, err
	}

	err = migration.MigrationDB(db)
	if err != nil {
		return nil, err
	}

	app := r.NewGin()

	return &Server{
		app: app,
		db:  db,
	}, nil
}

type Server struct {
	app        *gin.Engine
	db         *gorm.DbPgSql
	httpRouter *r.HTTPRouter
}

func (s *Server) Close() error {
	// closing all services
	s.db.Close()
	return nil
}

func (s *Server) InitializeServices() error {
	repository := module.NewRepository(s.db)
	usecase := module.NewUsecase(repository)

	s.httpRouter = r.NewHTTPRouter(s.app, usecase)
	s.httpRouter.InitRouters()

	return nil
}

func (s *Server) Run() {
	port := cfg.HTTP_PORT
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Addr:    cfg.IP_ADDRESS + ":" + port,
		Handler: s.app,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		err := srv.ListenAndServe()
		if err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	log.Println("--------------- Receive Shutdown Signal ---------------")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// closing services...
	_ = s.Close()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("--------------- Closing Services Complete ---------------")
}
