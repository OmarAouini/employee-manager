package api

import (
	"fmt"
	"log"
	"os"

	"github.com/OmarAouini/employee-manager/constants"
	"github.com/OmarAouini/employee-manager/core"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Server struct {
	Core  core.Core
	Fiber *fiber.App
}

func (s *Server) Init() {
	log.Printf("init server...\n")

	s.Fiber = fiber.New()
	s.Fiber.Use(logger.New())
	s.Fiber.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Content-Type, Authorization",
		AllowMethods: "GET, HEAD, OPTIONS, PUT, PATCH, POST, DELETE",
	}))
	s.Fiber.Use(recover.New())

	//routes init
	s.routes()
}

func (s *Server) Run() {
	s.Fiber.Listen(fmt.Sprintf("%s:%s", constants.HOST, constants.PORT))
}

func (s *Server) routes() {
	log.Printf("init api routes...\n\n")

	//log requests
	if os.Getenv("APP_ENV") == "local" {
		s.Fiber.Use(LogRequest())
	}

	//health
	s.Fiber.Get("/health", s.Health)

	//api
	api := s.Fiber.Group("/api")
	v1 := api.Group("/v1")

	companies := v1.Group("/companies")
	companies.Get("/", s.GetCompanies)

}
