-- +goose Up
CREATE TABLE IF NOT EXISTS calendars (
  id uuid PRIMARY KEY,
  title text NOT NULL,
  description text,
  location text,
  accept_responses_until timestamptz,
  password text,
  created_at timestamptz DEFAULT now() NOT NULL,
  updated_at timestamptz DEFAULT now() NOT NULL
);

CREATE TABLE IF NOT EXISTS calendar_time_slots (
  id uuid PRIMARY KEY,
  calendar_id uuid REFERENCES calendars(id) ON DELETE CASCADE,
  start_date timestamptz NOT NULL,
  end_date timestamptz NOT NULL,
  created_at timestamptz DEFAULT now() NOT NULL,
  updated_at timestamptz DEFAULT now() NOT NULL
);
CREATE INDEX idx_calendar_time_slots_calendar_id ON calendar_time_slots(calendar_id);

CREATE TABLE IF NOT EXISTS votes (
  id uuid PRIMARY KEY,
  calendar_id uuid REFERENCES calendars(id) ON DELETE CASCADE,
  calendar_time_slot_id uuid REFERENCES calendar_time_slots(id) ON DELETE CASCADE,
  username text NOT NULL,
  created_at timestamptz DEFAULT now() NOT NULL,
  updated_at timestamptz DEFAULT now() NOT NULL
);
CREATE INDEX idx_votes_calendar_id ON votes(calendar_id);
CREATE INDEX idx_votes_slot_id ON votes(calendar_time_slot_id);
CREATE UNIQUE INDEX idx_votes_user_slot ON votes(calendar_time_slot_id, username);

-- +goose Down
DROP TABLE IF EXISTS votes;
DROP TABLE IF EXISTS calendar_time_slots;
DROP TABLE IF EXISTS calendars;
