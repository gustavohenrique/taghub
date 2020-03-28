package domain

import (
	"time"
)

type Repo struct {
	ID          string     `json:"id,omitempty" db:"id"`
	CreatedAt   *time.Time `json:"created_at,omitempty" db:"created_at"`
	Name        string     `json:"name,omitempty" db:"name"`
	Description string     `json:"description,omitempty" db:"description"`
	URL         string     `json:"url,omitempty" db:"url"`
	Homepage    string     `json:"homepage,omitempty" db:"homepage"`
	Tags        []Tag      `json:"tags" db:"-"`
}

type Tag struct {
	ID         string `json:"id" db:"id"`
	Name       string `json:"name,omitempty" db:"name"`
	TotalRepos int    `json:"total_repos,omitempty" db:"-"`
}
