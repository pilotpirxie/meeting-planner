-- name: CreateCalendar :one
INSERT INTO calendars (
  id,
  title,
  description,
  location,
  accept_responses_until,
  created_at,
  updated_at
)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetCalendarByID :one
SELECT id, title, description, location, accept_responses_until, created_at, updated_at
FROM calendars
WHERE id = $1;

-- name: DeleteCalendarByID :exec
DELETE FROM calendars
WHERE id = $1;