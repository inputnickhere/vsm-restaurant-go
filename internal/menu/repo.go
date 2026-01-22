package menu

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

func (r *Repo) Create(ctx context.Context, name string, price int, isActive bool) (Item, error) {
	var it Item
	err := r.pool.QueryRow(ctx, `
		INSERT INTO menu_items (name, price, is_active)
		VALUES ($1, $2, $3)
		RETURNING id, name, price, is_active, created_at, updated_at
	`, name, price, isActive).Scan(
		&it.ID, &it.Name, &it.Price, &it.IsActive, &it.CreatedAt, &it.UpdatedAt,
	)
	return it, err
}

func (r *Repo) Update(ctx context.Context, id int64, name string, price int, isActive bool) (Item, error) {
	var it Item
	err := r.pool.QueryRow(ctx, `
		UPDATE menu_items
		SET name=$2, price=$3, is_active=$4, updated_at=now()
		WHERE id=$1
		RETURNING id, name, price, is_active, created_at, updated_at
	`, id, name, price, isActive).Scan(
		&it.ID, &it.Name, &it.Price, &it.IsActive, &it.CreatedAt, &it.UpdatedAt,
	)
	return it, err
}

func (r *Repo) Delete(ctx context.Context, id int64) error {
	_, err := r.pool.Exec(ctx, `DELETE FROM menu_items WHERE id=$1`, id)
	return err
}

func (r *Repo) List(ctx context.Context) ([]Item, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT id, name, price, is_active, created_at, updated_at
		FROM menu_items
		ORDER BY id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make([]Item, 0)
	for rows.Next() {
		var it Item
		if err := rows.Scan(&it.ID, &it.Name, &it.Price, &it.IsActive, &it.CreatedAt, &it.UpdatedAt); err != nil {
			return nil, err
		}
		out = append(out, it)
	}
	return out, rows.Err()
}

func (r *Repo) ListActive(ctx context.Context) ([]Item, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT id, name, price, is_active, created_at, updated_at
		FROM menu_items
		WHERE is_active = true
		ORDER BY id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make([]Item, 0)
	for rows.Next() {
		var it Item
		if err := rows.Scan(&it.ID, &it.Name, &it.Price, &it.IsActive, &it.CreatedAt, &it.UpdatedAt); err != nil {
			return nil, err
		}
		out = append(out, it)
	}
	return out, rows.Err()
}
