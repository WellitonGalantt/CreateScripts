-- +goose Up
CREATE TABLE subscriptions (
    id                      UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id                 UUID NOT NULL,
    plan_id                 UUID NOT NULL,
    status                  VARCHAR(20) NOT NULL, -- 'active' | 'canceled' | 'expired' | 'trial'
    current_period_start    TIMESTAMP WITH TIME ZONE NOT NULL,
    current_period_end      TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at              TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    
    CONSTRAINT fk_sub_user
        FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    CONSTRAINT fk_sub_plan
        FOREIGN KEY (plan_id) REFERENCES plans (id) ON DELETE RESTRICT
);

CREATE INDEX idx_subscriptions_user ON subscriptions (user_id);
CREATE INDEX idx_subscriptions_status ON subscriptions (status);