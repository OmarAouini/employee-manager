package web

import (
	"encoding/json"
	"net/http"
)

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	emps := []string{"wella", "willa", "wulla"}
	_ = json.NewEncoder(w).Encode(&ApiResponse{Message: "OK", Data: emps})
}
