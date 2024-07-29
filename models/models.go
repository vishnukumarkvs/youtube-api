package models

import "time"

type Video struct {
	ID          string    `db:"id" json:"id"`
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	PublishedAt time.Time `db:"published_at" json:"published_at"`
	Thumbnail   string    `db:"thumbnail" json:"thumbnail"`
}
