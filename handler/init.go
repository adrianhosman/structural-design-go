package handler

import (
	"github.com/adrianhosman/structural-design-go/usecase"
)

type Handler struct {
	usecase usecase.Usecase
}

func New(usecase usecase.Usecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}
