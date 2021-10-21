package main

import (
	"context"
)

func main() {
	// Соединение с экземпляром Postgres
	ctx := context.Background()
	dbpool := psql.Connect(ctx)
	defer dbpool.Close()
}
