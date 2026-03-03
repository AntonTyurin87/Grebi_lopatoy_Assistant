-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS main_data (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    chat_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL UNIQUE,
    step INTEGER,
    is_created BOOL,
    profession TEXT,
    gender INTEGER,
    world TEXT,
    marriage BOOL,
    childs INTEGER,
    wishes INTEGER,
    turn INTEGER,
    salary INTEGER,
    salary_extra_name TEXT,
    salary_extra INTEGER,
    cost_house INTEGER,
    cost_food INTEGER,
    cost_transport INTEGER,
    cost_cloth INTEGER,
    cost_extra_name TEXT,
    cost_extra INTEGER,
    total_income INTEGER,
    total_outcome INTEGER,
    flow INTEGER,
    cash INTEGER,
    menu_id INTEGER,
    id_last_active INTEGER,
    debt INTEGER,
    created_at TEXT,
    updated_at TEXT
);

-- Индексы
CREATE INDEX IF NOT EXISTS idx_main_data_user_id ON main_data(user_id);
CREATE INDEX IF NOT EXISTS idx_main_data_created_at ON main_data(created_at);
CREATE INDEX IF NOT EXISTS idx_main_data_updated_at ON main_data(updated_at);

-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS main_data;

-- Удаление индексов
DROP INDEX IF EXISTS idx_main_data_updated_at;
DROP INDEX IF EXISTS idx_main_data_created_at;
DROP INDEX IF EXISTS idx_main_data_user_id;
-- +goose StatementEnd
