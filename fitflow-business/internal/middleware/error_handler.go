package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// ErrorResponse represents a standardized error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
	Code    int    `json:"code"`
}

// GlobalErrorHandler is a middleware that catches all errors and panics
// It processes errors after handlers complete and before response is sent
func GlobalErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Process request
		c.Next()

		// Check for errors set by handlers using c.Error()
		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			
			// Determine status code based on error type and message
			statusCode := determineStatusCode(err)
			
			// Get error message
			errorMsg := err.Error()
			
			// Return error response
			c.JSON(statusCode, ErrorResponse{
				Error: errorMsg,
				Code:  statusCode,
			})
			c.Abort()
			return
		}
	}
}

// CustomRecovery is a custom recovery middleware that catches panics
func CustomRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic recovered: %v", err)
				
				errorMsg := "Internal server error"
				if str, ok := err.(string); ok {
					errorMsg = str
				} else if e, ok := err.(error); ok {
					errorMsg = e.Error()
				}

				c.JSON(http.StatusInternalServerError, ErrorResponse{
					Error: errorMsg,
					Code:  http.StatusInternalServerError,
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}

// determineStatusCode determines the HTTP status code based on error type and message
func determineStatusCode(err *gin.Error) int {
	// Handle binding errors
	if err.Type == gin.ErrorTypeBind {
		return http.StatusBadRequest
	}
	
	// Handle public errors (errors set by handlers)
	if err.Type == gin.ErrorTypePublic {
		errorMsg := strings.ToLower(err.Error())
		
		// Not found errors
		if strings.Contains(errorMsg, "not found") ||
		   strings.Contains(errorMsg, "record not found") ||
		   strings.Contains(errorMsg, "does not exist") {
			return http.StatusNotFound
		}
		
		// Validation errors
		if strings.Contains(errorMsg, "required") ||
		   strings.Contains(errorMsg, "invalid") ||
		   strings.Contains(errorMsg, "cannot be empty") {
			return http.StatusBadRequest
		}
		
		// Conflict errors
		if strings.Contains(errorMsg, "already exists") ||
		   strings.Contains(errorMsg, "duplicate") {
			return http.StatusConflict
		}
		
		// Unauthorized errors
		if strings.Contains(errorMsg, "unauthorized") ||
		   strings.Contains(errorMsg, "forbidden") ||
		   strings.Contains(errorMsg, "permission") {
			return http.StatusUnauthorized
		}
	}
	
	// Default to internal server error
	return http.StatusInternalServerError
}

// HandleError is a helper function for handlers to use
// It adds an error to the context with appropriate type
func HandleError(c *gin.Context, err error, statusCode int) {
	if statusCode == 0 {
		statusCode = http.StatusInternalServerError
	}
	c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(statusCode)
	c.Abort()
}

