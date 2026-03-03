-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS financial_assets (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    main_id INTEGER NOT NULL,
    type INTEGER,
    ticker TEXT,
    cost INTEGER,
    quantity INTEGER,
    coupon INTEGER,
    time INTEGER,
    price INTEGER,
    rate INTEGER,
    FOREIGN KEY (main_id) REFERENCES main_data(id) ON DELETE CASCADE
    );

-- Индексы
CREATE INDEX IF NOT EXISTS idx_financial_assets_main_id ON financial_assets(main_id);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS financial_assets;

-- Удаление индексов
DROP INDEX IF EXISTS idx_financial_assets_main_id;
-- +goose StatementEnd
