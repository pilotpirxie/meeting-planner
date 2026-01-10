-- name: CreateVote :one
INSERT INTO votes (
  id,
  calendar_id,
  calendar_time_slot_id,
  username,
  available,
  created_at,
  updated_at
)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: ListVotesByCalendarID :many
SELECT id, calendar_id, username, available, created_at
FROM votes
WHERE calendar_id = $1
ORDER BY created_at ASC;

-- name: DeleteVotesByID :exec
DELETE FROM votes
WHERE id = $1;