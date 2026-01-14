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
	params := sqlc.CreateCalendarParams{
		Title:       input.Title,
		Description: input.Description,
		Location:    input.Location,
	}

	if input.AcceptResponsesUntil != nil {
		params.AcceptResponsesUntil = pgtype.Timestamptz{
			Time:  *input.AcceptResponsesUntil,
			Valid: true,
		}
	}

	calendarID, err := s.queries.CreateCalendar(ctx, params)
	if err != nil {
		return pgtype.UUID{}, fmt.Errorf("failed to create calendar: %w", err)
	}

	return calendarID, nil
}
