package ingredients

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repo struct {
	pool *pgxpool.Pool
}

func NewRepo(pool *pgxpool.Pool) *Repo {
	return &Repo{pool: pool}
}

func (r *Repo) Create(ctx context.Context, name string, stock int) (Ingredient, error) {
	var it Ingredient
	err := r.pool.QueryRow(ctx, `
		INSERT INTO ingredients (name, stock)
		VALUES ($1, $2)
		RETURNING id, name, stock, created_at, updated_at
	`, name, stock).Scan(&it.ID, &it.Name, &it.Stock, &it.CreatedAt, &it.UpdatedAt)
	return it, err
}

func (r *Repo) Update(ctx context.Context, id int64, name string, stock int) (Ingredient, error) {
	var it Ingredient
	err := r.pool.QueryRow(ctx, `
		UPDATE ingredients
		SET name=$2, stock=$3, updated_at=now()
		WHERE id=$1
		RETURNING id, name, stock, created_at, updated_at
	`, id, name, stock).Scan(&it.ID, &it.Name, &it.Stock, &it.CreatedAt, &it.UpdatedAt)
	return it, err
}

func (r *Repo) Delete(ctx context.Context, id int64) error {
	_, err := r.pool.Exec(ctx, `DELETE FROM ingredients WHERE id=$1`, id)
	return err
}

func (r *Repo) List(ctx context.Context) ([]Ingredient, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT id, name, stock, created_at, updated_at
		FROM ingredients
		ORDER BY id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make([]Ingredient, 0)
	for rows.Next() {
		var it Ingredient
		if err := rows.Scan(&it.ID, &it.Name, &it.Stock, &it.CreatedAt, &it.UpdatedAt); err != nil {
			return nil, err
		}
		out = append(out, it)
	}
	return out, rows.Err()
}

// Restock adds delta to stock (can be positive).
func (r *Repo) Restock(ctx context.Context, id int64, delta int) (Ingredient, error) {
	var it Ingredient
	err := r.pool.QueryRow(ctx, `
		UPDATE ingredients
		SET stock = stock + $2, updated_at=now()
		WHERE id=$1
		RETURNING id, name, stock, created_at, updated_at
	`, id, delta).Scan(&it.ID, &it.Name, &it.Stock, &it.CreatedAt, &it.UpdatedAt)
	return it, err
}
