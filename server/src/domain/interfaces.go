package domain

import (
	"server/libs/filter"
)

type RepoRepository interface {
	Create(item Repo) error
	Update(item Repo) error
	ReadOne(item Repo) (Repo, error)
	Search(req filter.Request) ([]Repo, int, error)
}

type RepoService interface {
	Create(item Repo) error
	Update(item Repo) error
	ReadOne(item Repo) (Repo, error)
	Search(req filter.Request) ([]Repo, int, error)
	Sync() ([]Repo, error)
}
