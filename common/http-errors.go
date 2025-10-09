package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Package level dispatch functions
func DispatchInternalError(c *gin.Context, err interface{}) {
	if err == nil {
		err = gin.H{}
	}
	c.JSON(http.StatusInternalServerError, Response{
		Success: false,
		Message: "something went wrong, please try again.",
		Data:    err,
	})
}

func DispatchBadRequest(c *gin.Context, msg interface{}, err error, errorType ...string) {
	c.JSON(http.StatusBadRequest, Response{
		Success: false,
		Message: msg,
		Type:    errorType,
		Data:    err,
	})
}

func DispatchNotFound(c *gin.Context, msg string, err error) {
	c.JSON(http.StatusNotFound, Response{
		Success: false,
		Message: msg,
		Data:    err,
	})
}

func DispatchForbidden(c *gin.Context, msg string) {
	c.JSON(http.StatusForbidden, Response{
		Success: false,
		Message: msg,
		Data:    nil,
	})
}

func DispatchMethodNotAllowed(c *gin.Context, msg string) {
	c.JSON(http.StatusMethodNotAllowed, Response{
		Success: false,
		Message: msg,
		Data:    nil,
	})
}

// Response structure used by all dispatchers
type Response struct {
	Success bool        `json:"success"`
	Message interface{} `json:"message"`
	Type    []string    `json:"type,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// ErrorBuilder for creating structured errors
type ErrorBuilder struct{}

func NewErrorBuilder() *ErrorBuilder {
	return &ErrorBuilder{}
}

func (b *ErrorBuilder) New(statusCode int, message string, success bool) *APIError {
	return &APIError{
		Message:    message,
		Success:    success,
		StatusCode: statusCode,
	}
}

func (b *ErrorBuilder) BadRequest(message string) *APIError {
	return b.New(http.StatusBadRequest, message, false)
}

func (b *ErrorBuilder) Unauthorized(message string) *APIError {
	return b.New(http.StatusUnauthorized, message, false)
}

func (b *ErrorBuilder) NotFound(message string) *APIError {
	return b.New(http.StatusNotFound, message, false)
}

func (b *ErrorBuilder) InternalServerError(message string) *APIError {
	return b.New(http.StatusInternalServerError, message, false)
}

func (b *ErrorBuilder) Forbidden(message string) *APIError {
	return b.New(http.StatusForbidden, message, false)
}

func (b *ErrorBuilder) TooManyRequests(message string) *APIError {
	return b.New(http.StatusTooManyRequests, message, false)
}

func (b *ErrorBuilder) Accepted(message string) *APIError {
	return b.New(http.StatusAccepted, message, true)
}

// APIError represents a structured API error
type APIError struct {
	Message    string `json:"message"`
	Success    bool   `json:"success"`
	StatusCode int    `json:"code"`
}

// Helper to send APIError directly to response
func (e *APIError) Send(c *gin.Context) {
	c.JSON(e.StatusCode, e)
}
