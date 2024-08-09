package postgresdb

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/zeze322/wt-guided-weaponry/models"
)

type Storage interface {
	Categories(context.Context) ([]models.Category, error)
}

type PostgresConn struct {
	conn *pgx.Conn
}

func New(ctx context.Context, postgresURL string) (*PostgresConn, error) {
	conn, err := pgx.Connect(ctx, postgresURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %s", err)
	}

	if err := conn.Ping(ctx); err != nil {
		return nil, err
	}

	return &PostgresConn{
		conn: conn,
	}, nil
}

func (p *PostgresConn) Close(ctx context.Context) error {
	return p.conn.Close(ctx)
}

func (p *PostgresConn) Categories(ctx context.Context) ([]models.Category, error) {
	query := `SELECT category FROM categories`

	rows, err := p.conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var categories []models.Category

	for rows.Next() {
		category := models.Category{}
		if err := rows.Scan(&category.Name); err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("failed to get categories")
	}

	return categories, nil
}
