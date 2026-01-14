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

func (h *Handler) CreateCalendarEndpoint(w http.ResponseWriter, r *http.Request) {
	var requestBody CreateCalendarRequest

	if parsingError := ParseRequest(r, RequestOptions{Body: &requestBody}); parsingError != nil {
		RespondError(w, http.StatusBadRequest, parsingError.Error())
		return
	}

	serviceInput := services.CreateCalendarInput{
		Title:       requestBody.Title,
		Description: requestBody.Description,
		Location:    requestBody.Location,
	}

	if requestBody.AcceptResponsesUntil != nil {
		parsedTime, timeParsingError := time.Parse(time.RFC3339, *requestBody.AcceptResponsesUntil)
		if timeParsingError != nil {
			RespondError(w, http.StatusBadRequest, "Invalid time format for accept_responses_until, expected RFC3339")
			return
		}
		serviceInput.AcceptResponsesUntil = &parsedTime
	}

	calendarID, creationError := h.CalendarService.CreateCalendar(r.Context(), serviceInput)
	if creationError != nil {
		RespondError(w, http.StatusInternalServerError, "Failed to create calendar")
		return
	}

	RespondJSON(w, http.StatusCreated, CreateCalendarResponse{
		ID: utils.UUIDToString(calendarID),
	})
}
