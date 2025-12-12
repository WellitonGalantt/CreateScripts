-- +goose Up
CREATE TABLE user_points (
    user_id     UUID PRIMARY KEY,  -- 1 registro por usuário
    points      INTEGER NOT NULL DEFAULT 0,
    updated_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    
    CONSTRAINT fk_user_points_user
        FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);