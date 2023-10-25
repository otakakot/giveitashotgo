package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func main() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := conn.Close(context.Background()); err != nil {
			slog.Warn(err.Error())
		}
	}()

	var res Test

	ins := "INSERT INTO tests DEFAULT VALUES RETURNING id, created_at, updated_at, deleted"

	if err := conn.QueryRow(context.Background(), ins).Scan(&res.ID, &res.CreatedAt, &res.UpdatedAt, &res.Deleted); err != nil {
		panic(err)
	}

	slog.Info(fmt.Sprintf("INSERT: %+v", res))

	upd := "UPDATE tests SET deleted = TRUE WHERE id = $1 RETURNING id, created_at, updated_at, deleted"

	if err := conn.QueryRow(context.Background(), upd, res.ID).Scan(&res.ID, &res.CreatedAt, &res.UpdatedAt, &res.Deleted); err != nil {
		panic(err)
	}

	slog.Info(fmt.Sprintf("UPDATE: %+v", res))
}

type Test struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Deleted   bool
}
