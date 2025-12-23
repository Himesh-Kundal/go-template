package handlers

import (
	"go-template/internal/db"
)

type Handler struct {
	q *db.Queries
}

func New(q *db.Queries) *Handler {
	return &Handler{q: q}
}
