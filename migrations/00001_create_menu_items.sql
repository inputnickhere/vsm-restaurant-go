
-- +goose Up
CREATE TABLE menu_items (
    id          BIGSERIAL PRIMARY KEY,
    name        TEXT NOT NULL,
    price       INTEGER NOT NULL,
    is_active   BOOLEAN NOT NULL DEFAULT TRUE,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX menu_items_is_active_idx ON menu_items (is_active);


-- +goose Down
DROP TABLE menu_items;
