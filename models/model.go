package models


//Article ...
type Article struct {
	ID        int64    ` json:"id"`
	UserID    int64    ` json:"-"`
	Title     string   ` json:"title"`
	Content   string   ` json:"content"`
	UpdatedAt int64    ` json:"updated_at"`
	CreatedAt int64    ` json:"created_at"`
	User      *Thing ` json:"user"`
}