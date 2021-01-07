package internal

type User struct {
	ID        int
	Email     string
	Username  string
	CreatedAt time.time
}

type TodoItem struct {
	ID          int
	Content     string
	UserID      int
	CreatedAt   time.time
	CompletedAt time.time
}
