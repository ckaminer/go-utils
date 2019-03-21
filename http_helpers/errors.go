package http_helpers

import (
	"fmt"
	"net/http"
)

type HttpError struct {
	StatusCode int
	Message    string
}

func (e HttpError) Error() string {
	return e.Message
}

type ResourceError interface {
	Error() string
	Code() int
	ClientMessage() string
}

type NotFoundError struct {
	Message    string
	EntityType string
}

func (e NotFoundError) Error() string {
	return e.Message
}

func (e NotFoundError) Code() int {
	return http.StatusNotFound
}

func (e NotFoundError) ClientMessage() string {
	return fmt.Sprintf("%v not found", e.EntityType)
}

type ConflictError struct {
	Message    string
	EntityType string
}

func (e ConflictError) Error() string {
	return e.Message
}

func (e ConflictError) Code() int {
	return http.StatusConflict
}

func (e ConflictError) ClientMessage() string {
	return fmt.Sprintf("Unable to create %v: conflict", e.EntityType)
}
