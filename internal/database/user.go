package database

import (
	"context"
	"database/sql"

	"github.com/trongbq/gotodo-server/internal"
)

// DBUser is a wrapper of internal.User,
// to prevent exposing sensitive data of user data like password, ...
type DBUser struct {
	internal.User
	Password string
}

func (db *DB) GetUser(ctx context.Context, id int64) (*internal.User, error) {
	row := db.QueryRow(ctx, "SELECT id, email, username, created_at FROM user WHERE id = ?", id)
	user, err := scan(row.Scan)
	if err == sql.ErrNoRows {
		return nil, ErrNoRecordFound
	}
	return user, err
}

func (db *DB) InsertUser(ctx context.Context, user DBUser) (int64, error) {
	r, err := db.Exec(ctx, "INSERT INTO user(email, username, password, created_at) values(?, ?, ?, ?)", user.Email, user.Username, user.Password, user.CreatedAt)
	if err != nil {
		return 0, err
	}
	return r.LastInsertId()
}

func (db *DB) GetUserIDAndPasswordByUsername(ctx context.Context, username string) (int64, string, error) {
	row := db.QueryRow(ctx, "SELECT id, password FROM user WHERE username = ?", username)
	var password string
	var id int64
	err := row.Scan(&id, &password)
	return id, password, err
}

func scan(scan func(dest ...interface{}) error) (*internal.User, error) {
	var u internal.User
	err := scan(&u.ID, &u.Email, &u.Username, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
