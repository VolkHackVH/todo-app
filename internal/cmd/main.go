package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/VolkHackVH/todo-list/internal/db"
	"github.com/VolkHackVH/todo-list/internal/router"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print(".env file not loaded", err)
	}

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, os.Getenv("POSTGRES_URI"))
	if err != nil {
		log.Fatal("failed to connection database", err)
	}
	defer conn.Close(ctx)

	fmt.Println("DATABASE CONNETCTED")
	queries := db.New(conn)

	r := router.InitRouter(queries)

	r.Run("localhost:8080")
}
