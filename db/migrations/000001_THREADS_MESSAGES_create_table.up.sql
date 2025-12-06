BEGIN;

CREATE TABLE IF NOT EXISTS threads (
    id uuid PRIMARY KEY,
    title TEXT NULL
);

CREATE TABLE IF NOT EXISTS messages (
    id uuid PRIMARY KEY,
    thread_id uuid REFERENCES threads(id),
    content TEXT NOT NULL
);

COMMIT;
