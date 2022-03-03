package apiresponse

//response handler wrapper
type ApiResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
