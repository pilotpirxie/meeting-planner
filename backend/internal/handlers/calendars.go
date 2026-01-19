package handlers

import (
	"meeting-planner/backend/internal/services"
	"meeting-planner/backend/internal/utils"
	"net/http"
	"time"
)

type CreateCalendarRequest struct {
	Title                string  `json:"title" validate:"required,min=3,max=256"`
	Description          *string `json:"description,omitempty" validate:"omitempty,max=1024"`
	Location             *string `json:"location,omitempty" validate:"omitempty,max=512"`
	AcceptResponsesUntil *string `json:"accept_responses_until,omitempty" validate:"omitempty,rfc3339"`
	Password 					   *string `json:"password,omitempty" validate:"omitempty,min=3,max=128"`
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

type CalendarTimeSlots struct {
	StartDate string `json:"start_date" validate:"required,rfc3339"`
	EndDate   string `json:"end_date" validate:"required,rfc3339"`
}

type CreateCalendarTimeSlotsRequest struct {
	TimeSlots []CalendarTimeSlots `json:"time_slots" validate:"required,dive,required"`
}

func (h *Handler) CreateCalendarTimeSlotsEndpoint(w http.ResponseWriter, r *http.Request) {
	calendarID := r.PathValue("calendar_id")
	
	var requestBody CreateCalendarTimeSlotsRequest

	if parsingError := ParseRequest(r, RequestOptions{Body: &requestBody}); parsingError != nil {
		RespondError(w, http.StatusBadRequest, parsingError.Error())
		return
	}

	calendarUUID, uuidError := utils.StringToUUID(calendarID)
	if uuidError != nil {
		RespondError(w, http.StatusBadRequest, "Invalid calendar ID")
		return
	}

	var timeSlots []services.TimeSlotInput
	for _, slot := range requestBody.TimeSlots {
		startTime, _ := time.Parse(time.RFC3339, slot.StartDate)
		endTime, _ := time.Parse(time.RFC3339, slot.EndDate)

		if !endTime.After(startTime) {
			RespondError(w, http.StatusBadRequest, "end_date must be after start_date")
			return
		}

		timeSlots = append(timeSlots, services.TimeSlotInput{
			StartDate: startTime,
			EndDate:   endTime,
		})
	}

	serviceInput := services.CreateCalendarTimeSlotsInput{
		CalendarID: calendarUUID,
		TimeSlots:  timeSlots,
	}

	creationError := h.CalendarService.CreateCalendarTimeSlots(r.Context(), serviceInput)
	if creationError != nil {
		RespondError(w, http.StatusInternalServerError, "Failed to create calendar time slots")
		return
	}

	w.WriteHeader(http.StatusCreated)
}
