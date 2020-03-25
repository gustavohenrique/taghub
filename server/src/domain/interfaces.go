package domain

import (
	"server/libs/filter"
)

type RepoRepository interface {
	Create(item Repo) (Repo, error)
	Update(item Repo) (Repo, error)
	ReadOne(item Repo) (Repo, error)
	Search(req filter.Request) ([]Repo, int, error)
}

type RepoService interface {
	Create(item Repo) (Repo, error)
	Update(item Repo) (Repo, error)
	ReadOne(item Repo) (Repo, error)
	Search(req filter.Request) ([]Repo, int, error)
}
