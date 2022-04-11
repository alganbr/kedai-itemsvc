package databases

import (
	"context"
	"fmt"
	"github.com/alganbr/kedai-itemsvc/configs"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
)

type DB struct {
	Pool *pgxpool.Pool
}

func NewDB(cfg *configs.Config) *DB {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.Database.Item,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
	)

	pool, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		panic(err)
	}

	return &DB{
		Pool: pool,
	}
}

func (d *DB) RunMigration(cfg *configs.Config) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Database.Item,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
	)

	m, err := migrate.New(cfg.Database.Migration, connString)
	defer m.Close()
	if err != nil {
		panic(err)
	}

	m.Up()
}
