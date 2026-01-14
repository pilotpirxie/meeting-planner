package handlers

import (
	"meeting-planner/backend/internal/services"
	"meeting-planner/backend/internal/utils"
	"net/http"
	"time"
)

type CreateCalendarRequest struct {
	Title                string  `json:"title" validate:"required,min=3,max=200"`
	Description          *string `json:"description,omitempty" validate:"omitempty,max=1000"`
	Location             *string `json:"location,omitempty" validate:"omitempty,max=500"`
	AcceptResponsesUntil *string `json:"accept_responses_until,omitempty" validate:"omitempty,rfc3339"`
}

type CreateCalendarResponse struct {
	ID string `json:"id"`
}

func (h *Handler) CreateCalendar(w http.ResponseWriter, r *http.Request) {
	var req CreateCalendarRequest

	if err := ParseRequest(r, RequestOptions{Body: &req}); err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	input := services.CreateCalendarInput{
		Title:       req.Title,
		Description: req.Description,
		Location:    req.Location,
	}

	if req.AcceptResponsesUntil != nil {
		parsedTime, err := time.Parse(time.RFC3339, *req.AcceptResponsesUntil)
		if err != nil {
			RespondError(w, http.StatusBadRequest, "Invalid time format for accept_responses_until, expected RFC3339")
			return
		}
		input.AcceptResponsesUntil = &parsedTime
	}

	calendarID, err := h.CalendarService.CreateCalendar(r.Context(), input)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, "Failed to create calendar")
		return
	}

	RespondJSON(w, http.StatusCreated, CreateCalendarResponse{
		ID: utils.UUIDToString(calendarID),
	})
}
