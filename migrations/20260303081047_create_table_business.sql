-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS business (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    main_id INTEGER NOT NULL,
    type INTEGER,
    name TEXT,
    cost INTEGER,
    profit INTEGER,
    FOREIGN KEY (main_id) REFERENCES main_data(id) ON DELETE CASCADE
);

-- Индексы
CREATE INDEX IF NOT EXISTS idx_business_main_id ON business(main_id);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS business;

-- Удаление индексов
DROP INDEX IF EXISTS idx_business_main_id;
-- +goose StatementEnd
