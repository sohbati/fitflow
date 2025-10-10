package auth

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// SuccessResponse represents a successful response
type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
