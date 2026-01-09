BEGIN;

DROP INDEX IF EXISTS idx_messages_id;
DROP INDEX IF EXISTS idx_messages_created_at;
DROP INDEX IF EXISTS idx_messages_deleted_at;
DROP TABLE IF EXISTS messages;

DROP INDEX IF EXISTS idx_threads_id;
DROP INDEX IF EXISTS idx_threads_created_at;
DROP INDEX IF EXISTS idx_threads_deleted_at;
DROP TABLE IF EXISTS threads;

COMMIT;