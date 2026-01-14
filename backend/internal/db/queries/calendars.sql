-- name: CreateCalendar :one
INSERT INTO calendars (
  title,
  description,
  location,
  accept_responses_until
)
VALUES ($1, $2, $3, $4)
RETURNING id;

-- name: GetCalendarByID :one
SELECT *
FROM calendars
WHERE id = $1;

-- name: DeleteCalendarByID :exec
DELETE FROM calendars
WHERE id = $1;