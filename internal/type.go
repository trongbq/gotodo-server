package internal

import "time"

type User struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

type TodoItem struct {
	ID          int64
	Content     string
	UserID      int
	CreatedAt   time.Time
	CompletedAt time.Time
}
