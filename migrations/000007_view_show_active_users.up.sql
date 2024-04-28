BEGIN;

CREATE OR REPLACE VIEW view_active_users AS 
    SELECT * FROM users 
    where deleted_at IS NULL;

COMMIT;