package db

import (
	"time"

	"github.com/uptrace/bun"
)

type News struct {
	bun.BaseModel `bun:"table:news"`

	ID          int64     `bun:"id,pk,autoincrement"`
	Title       string    `bun:"title,notnull"`
	Description string    `bun:"description,notnull"`
	Category    string    `bun:"category,notnull"`
	CreatedAt   time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
}
