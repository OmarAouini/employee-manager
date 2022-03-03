package web

import (
	"encoding/json"
	"net/http"

	"github.com/OmarAouini/employee-manager/web/apiresponse"
)

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(&apiresponse.ApiResponse{Message: "OK", Data: "isOk"})
}
