CREATE TABLE IF NOT EXISTS role_has_permissions (
    id serial PRIMARY KEY,
    user_id BIGINT NOT NULL,
    permission_id BIGINT NOT NULL
);