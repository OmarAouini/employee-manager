package middlewares

import (
	"encoding/json"
	"net/http"

	"github.com/OmarAouini/employee-manager/web/apiresponse"
	"github.com/OmarAouini/employee-manager/web/auth"
)

func Protected(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := auth.ValidateToken(r.Header.Get("Authorization"))
		if err != nil {
			r.Response.StatusCode = 401
			_ = json.NewEncoder(w).Encode(apiresponse.ApiResponse{})
		}
		h.ServeHTTP(w, r)
	})
}
