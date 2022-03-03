package web

import (
	"fmt"
	"net/http"

	"github.com/OmarAouini/employee-manager/constants"
	"github.com/OmarAouini/employee-manager/web/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/sirupsen/logrus"
)

func RunServer() {
	r := chi.NewRouter()
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	r.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: logger}))
	r.Use(middleware.SetHeader("Content-Type", "application/json"))
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins: []string{"https://*", "http://*", },
		AllowedOrigins: []string{"*"}, //TODO: temp
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))
	configRoutes(r)
	http.ListenAndServe(fmt.Sprintf("%s:%v", constants.HOST, constants.PORT), r)
}

func configRoutes(r *chi.Mux) {
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		r.Response.StatusCode = 202
	})
	r.Get("/employees", middlewares.Protected(GetAllEmployees))
}
