-- +goose Up
CREATE TABLE menu_item_ingredients (
    menu_item_id   BIGINT NOT NULL REFERENCES menu_items(id) ON DELETE CASCADE,
    ingredient_id  BIGINT NOT NULL REFERENCES ingredients(id) ON DELETE RESTRICT,
    quantity       INTEGER NOT NULL,
    PRIMARY KEY (menu_item_id, ingredient_id)
);

CREATE INDEX mii_menu_item_id_idx ON menu_item_ingredients(menu_item_id);
CREATE INDEX mii_ingredient_id_idx ON menu_item_ingredients(ingredient_id);

-- +goose Down
DROP TABLE menu_item_ingredients;
