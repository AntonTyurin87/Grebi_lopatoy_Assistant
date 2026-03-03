-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS movable_property (
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
CREATE INDEX IF NOT EXISTS idx_movable_property_main_id ON movable_property(main_id);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS movable_property;

-- Удаление индексов
DROP INDEX IF EXISTS idx_movable_property_main_id;
-- +goose StatementEnd

