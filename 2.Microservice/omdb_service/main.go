package main

import (
	"fmt"
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"omdb_service/infrastructure/config"
	grpcinfra "omdb_service/infrastructure/grpc"
	"omdb_service/movie"
)

func init() {
	config.Init()
}

func main() {
	r := chi.NewRouter()
	setCors(r)
	r.Use(middleware.Logger)

	go InjectGrpc(config.Config)

	// inject movie routes
	movieRoutes := movie.NewRoute(config.Config)
	movieRoutes.RegisterRoutes(r)

	logrus.Infof("%s running on port %s\n", config.Config.ServiceName, config.Config.Server.Port)
	http.ListenAndServe(":"+config.Config.Server.Port, r)
}

func InjectGrpc(config config.AppConfig) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config.Server.GrpcPort))
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	movieGrpcServer := movie.NewGrpc(config)
	grpcinfra.RegisterMovieServer(s, movieGrpcServer)

	logrus.Infof("GRPC running on port %s\n", config.Server.GrpcPort)
	if err := s.Serve(lis); err != nil {
		logrus.Fatalf("failed to serve: %v", err)
	}
}

func setCors(r *chi.Mux) {
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
}
