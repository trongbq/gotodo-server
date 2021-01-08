package database

import (
	"context"
	"database/sql"

	"github.com/trongbq/gotodo-server/internal"
)

func (db *DB) GetUser(ctx context.Context, id int) (*internal.User, error) {
	row := db.QueryRow(ctx, "SELECT * FROM user WHERE id = ?", id)
	user, err := scan(row.Scan)
	if err == sql.ErrNoRows {
		return nil, ErrNoRecordFound
	}
	return user, err
}

func (db *DB) InsertUser(ctx context.Context, user internal.User) (int64, error) {
	r, err := db.Exec(ctx, "INSERT INTO user(email, username, created_at) values(?, ?, ?)", user.Email, user.Username, user.CreatedAt)
	if err != nil {
		return 0, err
	}
	return r.LastInsertId()
}

func scan(scan func(dest ...interface{}) error) (*internal.User, error) {
	var u internal.User
	err := scan(&u.ID, &u.Username, &u.Email, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
