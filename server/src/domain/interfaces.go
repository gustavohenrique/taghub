package domain

import (
	"server/libs/filter"
)

type TagRepository interface {
	Create(item Tag) (Tag, error)
}

type RepoRepository interface {
	Create(item Repo) error
	ReadOne(item Repo) (Repo, error)
	Search(req filter.Request) ([]Repo, int, error)
	SearchTag(filters filter.Request) ([]Tag, int, error)
	GetTagByName(name string) (Tag, error)
	RemoveTagFromRepo(repo Repo, tag Tag) error
	AddTagToRepo(repo Repo, tag Tag) error
}

type RepoService interface {
	Create(item Repo) error
	ReadOne(item Repo) (Repo, error)
	Search(req filter.Request) ([]Repo, int, error)
	GetTotalStarredRepositories() (int64, error)
	Sync() ([]Repo, int64, error)
	SearchTag(filters filter.Request) ([]Tag, int, error)
	RemoveTagFromRepo(repo Repo, tag Tag) error
	AddTagToRepo(repo Repo, tag Tag) (Tag, error)
}
