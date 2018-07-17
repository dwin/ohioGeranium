package model

import (
	"log"

	"github.com/jackc/pgx"
	_ "github.com/lib/pq" // postgres driver
)

const pgConnStr = "postgres://postgres:IpcUUrehVEE4MOJNpLKU@postgres/postgres?sslmode=disable"

var db = openDB()

// openDB starts postgres database connection pool
func openDB() *pgx.ConnPool {
	connConfig, err := pgx.ParseURI(pgConnStr)
	if err != nil {
		log.Fatalf("Unable to parse connection string error: %s", err)
	}
	p := pgx.ConnPoolConfig{ConnConfig: connConfig, MaxConnections: 100}
	db, err := pgx.NewConnPool(p)
	if err != nil {
		log.Fatalf("Unable to open postgres connection error: %s", err)
	}

	// Create Tables
	createTables(db)

	return db
}

// createTables creates tables required by application if not existing
func createTables(db *pgx.ConnPool) {
	_, err := db.Exec(`CREATE TABLE "public"."providers" (
		"provider_number" int NOT NULL,
		"data" jsonb NOT NULL,
		"created_at" timestamp without time zone,
		"updated_at" timestamp without time zone,
		"deleted_at" timestamp without time zone,
		PRIMARY KEY ("provider_number"),
		UNIQUE ("provider_number")
	);`)
	if err != nil {
		log.Printf("Unable to create table providers error: %s", err)
	}
	_, err = db.Exec(`CREATE TABLE "public"."queries" (
		"id" serial,
		"data" jsonb NOT NULL,
		"created_at" timestamp,
		"updated_at" timestamp,
		"deleted_at" timestamp,
		PRIMARY KEY ("id")
	);`)
	if err != nil {
		log.Printf("Unable to create table queries error: %s", err)
	}
}
