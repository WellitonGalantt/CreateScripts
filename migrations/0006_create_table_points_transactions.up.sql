-- +goose Up
CREATE TABLE points_transactions (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id     UUID NOT NULL,
    type        VARCHAR(20) NOT NULL,    -- 'credit' | 'debit'
    amount      INTEGER NOT NULL,        -- quantidade de pontos (+/-)
    reason      VARCHAR(50) NOT NULL,    -- 'generation' | 'batch' | 'manual' | 'monthly_reset' | etc.
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    
    CONSTRAINT fk_pt_user
        FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE INDEX idx_pt_user ON points_transactions (user_id);
CREATE INDEX idx_pt_created_at ON points_transactions (created_at);