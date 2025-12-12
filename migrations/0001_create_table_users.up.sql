-- +goose Up
CREATE TABLE users (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            VARCHAR(100) NOT NULL,
    email           VARCHAR(150) NOT NULL UNIQUE,
    password_hash   VARCHAR(255) NOT NULL,
    role            VARCHAR(20) NOT NULL DEFAULT 'user',  -- 'user' | 'admin'
    niche           VARCHAR(100),                         -- nicho principal
    goal            VARCHAR(255),                         -- objetivo (ex: crescer, perder vergonha...)
    created_at      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);