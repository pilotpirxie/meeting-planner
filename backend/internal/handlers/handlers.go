package handlers

import "meeting-planner/backend/internal/db"

type Handler struct {
	DB *db.DB
}

func New(db *db.DB) *Handler {
	return &Handler{
		DB: db,
	}
}
