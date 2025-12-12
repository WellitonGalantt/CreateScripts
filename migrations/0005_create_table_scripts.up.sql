-- +goose Up
CREATE TABLE scripts (
    id                  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id             UUID NOT NULL,
    type                VARCHAR(20) NOT NULL,      -- 'short' | 'long' | 'batch'
    style               VARCHAR(50),               -- emocional, engraçado, etc.
    topic               VARCHAR(255),              -- tema informado
    content             TEXT NOT NULL,             -- roteiro completo
    title               VARCHAR(255),              -- título sugerido
    description         TEXT,                      -- descrição gerada
    hashtags            TEXT,                      -- hashtags geradas
    thumbnail_prompt    TEXT,                      -- prompt da thumb
    is_favorite         BOOLEAN NOT NULL DEFAULT FALSE,
    created_at          TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    
    CONSTRAINT fk_scripts_user
        FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE INDEX idx_scripts_user ON scripts (user_id);
CREATE INDEX idx_scripts_created_at ON scripts (created_at);
CREATE INDEX idx_scripts_favorite ON scripts (user_id, is_favorite);