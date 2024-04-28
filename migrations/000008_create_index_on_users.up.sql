BEGIN;
CREATE INDEX deleted_at_idx ON users (deleted_at);
COMMIT;