package todo

import "time"

type Todo struct {
	ID          string
	Title       string
	Description string
	CreatedAt   time.Time
	Completed   bool
}
