package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/masterraf21/ecommerce-backend/configs"
	"github.com/masterraf21/ecommerce-backend/utils/mongodb"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/rs/cors"
)

// Server represents server
type Server struct {
	Instance    *mongo.Database
	Port        string
	ServerReady chan bool
}

func configureMongo() *mongo.Database {
	option := mongodb.Option{
		Hosts:    configs.MongoDB.Hosts,
		Database: configs.MongoDB.Database,
		Options:  configs.MongoDB.Options,
	}

	instance, err := mongodb.Init(option)
	if err != nil {
		log.Fatalf("%s: %s", "Failed to connect mongodb", err)
	}

	log.Println("MongoDB connection is successfully established!")

	return instance
}

// Start will start server
func (s *Server) Start() {
	port := configs.Server.Port
	if port == "" {
		port = "8000"
	}

	r := new(mux.Router)

	handler := cors.Default().Handler(r)

	srv := &http.Server{
		Handler:      handler,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		log.Printf("Starting server on port %s!", port)
		if err := srv.ListenAndServe(); err != nil {
			log.Println("Shutting Down Server...")
			log.Fatal(err.Error())
		}
	}()

	if s.ServerReady != nil {
		s.ServerReady <- true
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("failed to gracefully shutdown the server: %s", err)
	}
}