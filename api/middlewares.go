package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

//used to log the request and the response with the body, headers etc...
func LogRequestResponse() fiber.Handler {
	return func(c *fiber.Ctx) error {
		fmt.Println("\n=============================================\nREQUEST: " + c.Request().String())
		con := c.Next()
		fmt.Println("\n=============================================\nRESPONSE: " + c.Response().String())
		return con
	}
}

//used to log the requests only
func LogRequest() fiber.Handler {
	return func(c *fiber.Ctx) error {
		fmt.Println("\n=============================================\nREQUEST: " + c.Request().String() + "\n")
		return c.Next()
	}
}
