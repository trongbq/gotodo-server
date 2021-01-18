package database

import (
	"context"
	"time"

	"github.com/trongbq/gotodo-server/internal"
)

func (db *DB) GetTodoCountByUser(ctx context.Context, userID int64) (int64, error) {
	var count int64
	err := db.QueryRow(ctx, "SELECT COUNT(*) FROM todo WHERE user_id = ?", userID).Scan(&count)
	return count, err
}
func (db *DB) GetTodosByUser(ctx context.Context, userID int64, offset, limit int) ([]*internal.Todo, error) {
	rows, err := db.Query(ctx, "SELECT * FROM todo WHERE user_id = ? LIMIT ? OFFSET ?", userID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var todos []*internal.Todo
	for rows.Next() {
		todo, err := scanTodo(rows.Scan)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (db *DB) InsertTodo(ctx context.Context, t internal.Todo) (int64, error) {
	r, err := db.Exec(ctx, `INSERT INTO todo(
		content,
		user_id,
		created_at,
		updated_at
	) VALUES(?,?,?,?)`, t.Content, t.UserID, t.CreatedAt, t.UpdatedAt)
	if err != nil {
		return 0, err
	}
	return r.LastInsertId()
}

func (db *DB) UpdateTodoContent(ctx context.Context, id int64, c string) error {
	_, err := db.Exec(ctx, "UPDATE todo SET content = ?, updated_at = ? WHERE id = ?", c, time.Now(), id)
	return err
}

func (db *DB) UpdateTodoComplete(ctx context.Context, id int64) error {
	_, err := db.Exec(ctx, "UPDATE todo SET completed_at = ?, updated_at = ? WHERE id = ?", time.Now(), time.Now(), id)
	return err
}

func (db *DB) DeleteTodo(ctx context.Context, id int64) error {
	_, err := db.Exec(ctx, "DELETE FROM todo WHERE id = ?", id)
	return err
}

func scanTodo(scan func(dest ...interface{}) error) (*internal.Todo, error) {
	var t internal.Todo
	err := scan(&t.ID, &t.Content, &t.UserID, &t.CreatedAt, &t.CompletedAt, &t.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &t, nil
}
