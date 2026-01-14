package handlers

import (
	"meeting-planner/backend/internal/db"
	"meeting-planner/backend/internal/services"
)

type Handler struct {
	DB              *db.DB
	CalendarService *services.CalendarService
}

func New(db *db.DB) *Handler {
	return &Handler{
		DB:              db,
		CalendarService: services.NewCalendarService(db.Queries),
	}
}
