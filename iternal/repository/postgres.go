package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Ошибка загрузки файла .env:", err)
	}
}

type PostgresBase struct {
	db *bun.DB
}

func NewPostgresBase() *PostgresBase {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(sqldb, pgdialect.New())
	log.Println("Подключение к бд установлено!")
	return &PostgresBase{db: db}
}
func (d *PostgresBase) ResetModel(ctx context.Context, models ...interface{}) error {
	return d.db.ResetModel(ctx, models...)
}
