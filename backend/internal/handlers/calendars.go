package handlers

import (
	"fmt"
	"meeting-planner/backend/internal/db/sqlc"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func (h *Handler) CreateCalendar(w http.ResponseWriter, r *http.Request) {
	var payloadBody struct {
		Title                string  `json:"title"`
		Description          *string `json:"description,omitempty"`
		Location             *string `json:"location,omitempty"`
		AcceptResponsesUntil *string `json:"accept_responses_until,omitempty"`
	}

	err := ParseRequest(r, RequestOptions{
		Body: &payloadBody,
	})
	if err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	params := sqlc.CreateCalendarParams{
		Title:       payloadBody.Title,
		Description: payloadBody.Description,
		Location:    payloadBody.Location,
	}

	if payloadBody.AcceptResponsesUntil != nil {
		parsedTime, err := time.Parse(time.RFC3339, *payloadBody.AcceptResponsesUntil)
		if err != nil {
			RespondError(w, http.StatusBadRequest, "Invalid time format for AcceptResponsesUntil")
			return
		}

		params.AcceptResponsesUntil = pgtype.Timestamptz{
			Time:  parsedTime,
			Valid: true,
		}
	}

	calendarID, err := h.DB.Queries.CreateCalendar(r.Context(), params)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, fmt.Errorf("Failed to create calendar: %w", err).Error())
		return
	}

	RespondJSON(w, http.StatusCreated, map[string]any{
		"id": calendarID,
	})
}
