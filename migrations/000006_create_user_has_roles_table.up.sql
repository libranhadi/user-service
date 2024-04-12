CREATE TABLE IF NOT EXISTS user_has_roles (
    id serial PRIMARY KEY,
    user_id BIGINT NOT NULL,
    role_id BIGINT NOT NULL
);