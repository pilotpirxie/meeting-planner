-- +goose Up
ALTER TABLE calendars
  ALTER COLUMN id SET DEFAULT gen_random_uuid();

ALTER TABLE calendar_time_slots
  ALTER COLUMN id SET DEFAULT gen_random_uuid();

ALTER TABLE votes
  ALTER COLUMN id SET DEFAULT gen_random_uuid();

-- +goose Down
ALTER TABLE calendars
  ALTER COLUMN id DROP DEFAULT;

ALTER TABLE calendar_time_slots
  ALTER COLUMN id DROP DEFAULT;

ALTER TABLE votes
  ALTER COLUMN id DROP DEFAULT;
