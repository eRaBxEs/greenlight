package data

import (
	"time"
)

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`           // Add the omitempty directive
	Runtime   int32     `json:"runtime,omitempty,string"` // Add the omitempty directive
	Genres    []string  `json:"genres,omitempty"`         // Add the omitempty directive
	Version   int32     `json:"version"`
}
