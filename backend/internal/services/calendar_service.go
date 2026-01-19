package services

import (
	"context"
	"fmt"
	"meeting-planner/backend/internal/db/sqlc"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type CalendarService struct {
	queries *sqlc.Queries
}

func NewCalendarService(queries *sqlc.Queries) *CalendarService {
	return &CalendarService{
		queries: queries,
	}
}

type CreateCalendarInput struct {
	Title                string
	Description          *string
	Location             *string
	AcceptResponsesUntil *time.Time
}

func (s *CalendarService) CreateCalendar(ctx context.Context, input CreateCalendarInput) (pgtype.UUID, error) {
	queryParams := sqlc.CreateCalendarParams{
		Title:       input.Title,
		Description: input.Description,
		Location:    input.Location,
	}

	if input.AcceptResponsesUntil != nil {
		queryParams.AcceptResponsesUntil = pgtype.Timestamptz{
			Time:  *input.AcceptResponsesUntil,
			Valid: true,
		}
	}

	calendarID, creationError := s.queries.CreateCalendar(ctx, queryParams)
	if creationError != nil {
		return pgtype.UUID{}, fmt.Errorf("failed to create calendar: %w", creationError)
	}

	return calendarID, nil
}

type TimeSlotInput struct {
	StartDate time.Time
	EndDate   time.Time
}

type CreateCalendarTimeSlotsInput struct {
	CalendarID pgtype.UUID
	TimeSlots  []TimeSlotInput
}

func (s *CalendarService) CreateCalendarTimeSlots(ctx context.Context, input CreateCalendarTimeSlotsInput) error {
	for _, slot := range input.TimeSlots {
		queryParams := sqlc.CreateCalendarTimeSlotParams{
			CalendarID: input.CalendarID,
			StartDate: pgtype.Timestamptz{
				Time:  slot.StartDate,
				Valid: true,
			},
			EndDate: pgtype.Timestamptz{
				Time:  slot.EndDate,
				Valid: true,
			},
		}

		_, creationError := s.queries.CreateCalendarTimeSlot(ctx, queryParams)
		if creationError != nil {
			return fmt.Errorf("failed to create calendar time slot: %w", creationError)
		}
	}

	return nil
}
