BEGIN;

-- Create the THREADS table
-- we use a serial primary key for internal use and a uuid for external references
-- That id is automatically generated using gen_random_uuid()
CREATE TABLE IF NOT EXISTS threads (
    pk bigint generated always as identity PRIMARY KEY,
    id uuid UNIQUE NOT NULL DEFAULT gen_random_uuid(),
    title TEXT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX IF NOT EXISTS idx_threads_id ON threads(id);
CREATE INDEX IF NOT EXISTS idx_threads_created_at ON threads(created_at);

-- Create the MESSAGES table
CREATE TABLE IF NOT EXISTS messages (
    pk bigint generated always as identity PRIMARY KEY,
    id uuid UNIQUE NOT NULL DEFAULT gen_random_uuid(),
    thread_id uuid REFERENCES threads(id),
    content TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX IF NOT EXISTS idx_messages_id ON messages(id);
CREATE INDEX IF NOT EXISTS idx_messages_created_at ON messages(created_at);

COMMIT;
