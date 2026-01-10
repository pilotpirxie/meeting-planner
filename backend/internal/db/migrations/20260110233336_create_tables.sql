-- +goose Up
CREATE TABLE IF NOT EXISTS calendars (
  id uuid PRIMARY KEY,
  title text NOT NULL,
  description text,
  location text,
  accept_responses_until timestamptz,
  created_at timestamptz DEFAULT now() NOT NULL,
  updated_at timestamptz DEFAULT now() NOT NULL
);

CREATE TABLE IF NOT EXISTS calendar_time_slots (
  id uuid PRIMARY KEY,
  calendar_id uuid REFERENCES calendars(id) ON DELETE CASCADE,
  slot_date date NOT NULL,
  start_time time NOT NULL,
  end_time time NOT NULL,
  created_at timestamptz DEFAULT now() NOT NULL,
  updated_at timestamptz DEFAULT now() NOT NULL
);

CREATE TABLE IF NOT EXISTS votes (
  id uuid PRIMARY KEY,
  calendar_id uuid REFERENCES calendars(id) ON DELETE CASCADE,
  calendar_time_slot_id uuid REFERENCES calendar_time_slots(id) ON DELETE CASCADE,
  username text NOT NULL,
  available jsonb NOT NULL,
  created_at timestamptz DEFAULT now() NOT NULL,
  updated_at timestamptz DEFAULT now() NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS votes;
DROP TABLE IF EXISTS calendars;
DROP TABLE IF EXISTS calendar_time_slots;
