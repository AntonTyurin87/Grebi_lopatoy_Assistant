-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS real_estate (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    main_id INTEGER NOT NULL,
    type INTEGER,
    name TEXT,
    cost INTEGER,
    term INTEGER,
    payment INTEGER,
    FOREIGN KEY (main_id) REFERENCES main_data(id) ON DELETE CASCADE
    );

-- Индексы
CREATE INDEX IF NOT EXISTS idx_real_estate_main_id ON real_estate(main_id);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS real_estate;

-- Удаление индексов
DROP INDEX IF EXISTS idx_real_estate_main_id;
-- +goose StatementEnd
