CREATE TABLE IF NOT EXISTS subscriptions (
    id            SERIAL PRIMARY KEY,
    service_name  VARCHAR(255) NOT NULL,
    price         INTEGER NOT NULL CHECK (price >= 0),
    user_id       UUID NOT NULL,
    start_date    DATE NOT NULL,      -- будем хранить как первое число месяца: 2025-07-01
    end_date      DATE,               -- nullable
    created_at    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- Индексы для фильтрации по пользователю и сервису
CREATE INDEX idx_subscriptions_user_id ON subscriptions(user_id);
CREATE INDEX idx_subscriptions_service_name ON subscriptions(service_name);
