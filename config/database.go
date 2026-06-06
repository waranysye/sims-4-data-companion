package config

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func ConnectDB() *pgxpool.Pool {
	// Membaca konfigurasi dari environment variable (Docker), jika tidak ada pakai default localhost
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "28711" // Port luar untuk non-docker
	}
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "postgres"
	}
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "callmera" // ⚠️ Sesuaikan password lokalmu di sini
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "sims4_db"
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Gagal memproses konfigurasi database: %v\n", err)
		os.Exit(1)
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Gagal terhubung ke PostgreSQL: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("🚀 Sukses terhubung ke PostgreSQL (%s:%s) via Golang!\n", dbHost, dbPort)
	return pool
}
