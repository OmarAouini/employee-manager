package web

import (
	"encoding/json"
	"net/http"

	"github.com/OmarAouini/employee-manager/web/auth"
)

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDYzNDk0MDEsImlhdCI6MTY0NjM0MjIwMSwiaXNzIjoiZ29sYW5nLWVtcGxveWVlLWFwaSIsInN1YiI6InVzZXIifQ.v9W4XjEgxsZlIyQux0FfBYr9oDuzhtenNPfOsaZEmr8"
	err := auth.ValidateToken(tokenString)
	if err != nil {
		_ = json.NewEncoder(w).Encode(&ApiResponse{Message: "KO", Data: err})
	}
	claims, _ := auth.GetClaimsFromToken(tokenString)
	_ = json.NewEncoder(w).Encode(&ApiResponse{Message: "OK", Data: claims})
}
