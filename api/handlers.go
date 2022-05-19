package api

import (
	"github.com/gofiber/fiber/v2"
)

//api response wrapper
type apiResponse struct {
	Response string      `json:"response"`
	Message  *string     `json:"message,omitempty"`
	Data     interface{} `json:"data,omitempty"`
}

//used for ok case as return
func okReponse(message *string, data interface{}) apiResponse {
	return apiResponse{Response: "OK", Message: message, Data: data}
}

//used for error case as return
func errorReponse(message string) apiResponse {
	return apiResponse{Response: "KO", Message: &message}
}

//health check endpoint
func (s *Server) Health(c *fiber.Ctx) error {
	return c.Status(200).JSON(okReponse(nil, nil))
}

//COMPANIES
///////////////////////////////////

func (s *Server) GetCompanies(c *fiber.Ctx) error {
	cust, err := s.Core.GetCompanies()
	if err != nil {
		return c.Status(500).JSON(errorReponse(err.Error()))
	}
	return c.Status(200).JSON(okReponse(nil, cust))
}
