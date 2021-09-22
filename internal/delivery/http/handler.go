package http

import "github.com/jjmengze/mygo/internal/usecase"

// Handler http handler
type Handler struct {
	svc usecase.UseCase
}

// NewHandler create Handler instance
func NewHandler(svc usecase.UseCase) *Handler {
	return &Handler{
		svc: svc,
	}
}
