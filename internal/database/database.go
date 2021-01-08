package database

import (
	"context"
	"database/sql"
	"errors"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	ErrNoRecordFound = errors.New("No record found")
)

// DB type is a wrapper which provides external logging and others utilities for native sql.DB
type DB struct {
	db *sql.DB
}

func New(uri string) (*DB, error) {
	return open("mysql", uri)
}

func open(driverName string, uri string) (*DB, error) {
	db, err := sql.Open(driverName, uri)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

func (db *DB) QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return db.db.QueryRowContext(ctx, query, args...)
}

func (db *DB) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return db.db.ExecContext(ctx, query, args...)
}
