-- +goose Up
ALTER TABLE calendars
  ALTER COLUMN created_at SET DEFAULT now();

ALTER TABLE calendars
  ALTER COLUMN updated_at SET DEFAULT now();

ALTER TABLE calendar_time_slots
  ALTER COLUMN created_at SET DEFAULT now();

ALTER TABLE calendar_time_slots
  ALTER COLUMN updated_at SET DEFAULT now();

ALTER TABLE votes
  ALTER COLUMN created_at SET DEFAULT now();

ALTER TABLE votes
  ALTER COLUMN updated_at SET DEFAULT now();

-- +goose Down
ALTER TABLE calendars
  ALTER COLUMN created_at DROP DEFAULT;

ALTER TABLE calendars
  ALTER COLUMN updated_at DROP DEFAULT;

ALTER TABLE calendar_time_slots
  ALTER COLUMN created_at DROP DEFAULT;

ALTER TABLE calendar_time_slots
  ALTER COLUMN updated_at DROP DEFAULT;

ALTER TABLE votes 
  ALTER COLUMN created_at DROP DEFAULT;

ALTER TABLE votes 
  ALTER COLUMN updated_at DROP DEFAULT;