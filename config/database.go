package config

import (
	"context"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4/pgxpool"
)

func ConnectDB() *pgxpool.Pool {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "28711"
	}
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "postgres"
	}
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "callmera"
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "sims4_db"
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	// 🚀 JALANKAN MIGRASI OTOMATIS SEBELUM KONEK POOL
	migrationTarget := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
	m, err := migrate.New("file://db/migrations", migrationTarget)
	if err == nil {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			fmt.Printf("⚠️ Gagal menjalankan database migration: %v\n", err)
		} else {
			fmt.Println("🔄 Database migration berhasil disinkronkan!")
		}
	}

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		os.Exit(1)
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("🚀 Sukses terhubung ke PostgreSQL (%s:%s) via Golang!\n", dbHost, dbPort)
	return pool
}
