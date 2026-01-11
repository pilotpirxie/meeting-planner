-- name: GetCalendarTimeSlotsByCalendarID :many
SELECT 
  id,
  calendar_id,
  slot_date,
  start_time,
  end_time,
  created_at,
  updated_at
FROM calendar_time_slots
WHERE calendar_id = $1
ORDER BY slot_date, start_time;

-- name: CreateCalendarTimeSlot :one
INSERT INTO calendar_time_slots (
  id,
  calendar_id,
  slot_date,
  start_time,
  end_time,
  created_at,
  updated_at
)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: DeleteCalendarTimeSlotByID :exec
DELETE FROM calendar_time_slots
WHERE id = $1;
