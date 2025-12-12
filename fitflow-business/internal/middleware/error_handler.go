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

// GlobalErrorHandler is similar to Spring Boot's @ControllerAdvice
// It catches all errors from handlers and formats them consistently
func GlobalErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Process the request
		c.Next()

		// Check for errors set by handlers
		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			
			// Determine HTTP status code and format error response
			statusCode, errorKey := handleError(err)
			
			// Return standardized error response
			c.JSON(statusCode, ErrorResponse{
				Error: errorKey,
				Code:  statusCode,
			})
			c.Abort()
			return
		}
	}
}

// CustomRecovery handles panics (similar to Spring Boot's exception handling)
func CustomRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic recovered: %v", err)
				
				errorMsg := "internal_server_error"
				if str, ok := err.(string); ok {
					errorMsg = toUnderscoreKey(str)
				} else if e, ok := err.(error); ok {
					errorMsg = toUnderscoreKey(e.Error())
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

// handleError processes errors and determines status code and error key
// Similar to Spring Boot's exception handler methods
func handleError(err *gin.Error) (int, string) {
	errorMsg := err.Error()
	
	// Handle binding/validation errors
	if err.Type == gin.ErrorTypeBind {
		return http.StatusBadRequest, toUnderscoreKey(errorMsg)
	}
	
	// Handle public errors (errors from service layer)
	if err.Type == gin.ErrorTypePublic {
		errorKey := toUnderscoreKey(errorMsg)
		statusCode := determineStatusCode(errorKey)
		return statusCode, errorKey
	}
	
	// Default: internal server error
	return http.StatusInternalServerError, toUnderscoreKey(errorMsg)
}

// determineStatusCode determines HTTP status code based on error key
func determineStatusCode(errorKey string) int {
	errorKeyLower := strings.ToLower(errorKey)
	
	// Not found errors (404)
	if strings.Contains(errorKeyLower, "not_found") ||
		strings.Contains(errorKeyLower, "does_not_exist") {
		return http.StatusNotFound
	}
	
	// Validation errors (400)
	if strings.Contains(errorKeyLower, "required") ||
		strings.Contains(errorKeyLower, "invalid") ||
		strings.Contains(errorKeyLower, "cannot_be_empty") ||
		strings.Contains(errorKeyLower, "format") ||
		strings.Contains(errorKeyLower, "must_be") {
		return http.StatusBadRequest
	}
	
	// Conflict errors (409)
	if strings.Contains(errorKeyLower, "already_exists") ||
		strings.Contains(errorKeyLower, "duplicate") {
		return http.StatusConflict
	}
	
	// Unauthorized errors (401)
	if strings.Contains(errorKeyLower, "unauthorized") ||
		strings.Contains(errorKeyLower, "forbidden") ||
		strings.Contains(errorKeyLower, "permission") {
		return http.StatusUnauthorized
	}
	
	// Default: internal server error (500)
	return http.StatusInternalServerError
}

// toUnderscoreKey converts error message to underscore format
func toUnderscoreKey(message string) string {
	// If already in underscore format, return as-is
	if !strings.Contains(message, " ") {
		return message
	}
	
	// Convert to lowercase and replace spaces with underscores
	message = strings.ToLower(strings.TrimSpace(message))
	message = strings.ReplaceAll(message, " ", "_")
	
	// Remove multiple consecutive underscores
	for strings.Contains(message, "__") {
		message = strings.ReplaceAll(message, "__", "_")
	}
	
	// Remove leading/trailing underscores
	message = strings.Trim(message, "_")
	
	return message
}

// HandleError is a helper for handlers to add errors to context
// Handlers should use this instead of directly calling c.JSON for errors
func HandleError(c *gin.Context, err error, statusCode int) {
	if statusCode == 0 {
		statusCode = http.StatusInternalServerError
	}
	c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(statusCode)
	c.Abort()
}

