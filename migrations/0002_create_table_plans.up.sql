-- +goose Up
CREATE TABLE plans (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            VARCHAR(50) NOT NULL UNIQUE,   -- free, basic, pro, unlimited
    monthly_points  INTEGER NOT NULL DEFAULT 0,
    is_unlimited    BOOLEAN NOT NULL DEFAULT FALSE,
    price_cents     INTEGER NOT NULL DEFAULT 0,    -- preço em centavos (ex: 1990 = R$ 19,90)
    created_at      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);