package orm

import "time"

type Post struct {
	Id        int
	Title     string
	Content   string
	Author    string `sql:"not null"`
	CreatedAt time.Time
}
