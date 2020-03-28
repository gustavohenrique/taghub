package domain

import (
	"server/libs/filter"
)

type TagRepository interface {
	Create(item Tag) (Tag, error)
	ReadOne(item Tag) (Tag, error)
	Search(req filter.Request) ([]Tag, int, error)
	ReadAll() ([]Tag, error)
	GetTotalReposByTag(item Tag) (int, error)
    Remove(item Tag) error
    Update(item Tag) error
}

type TagService interface {
	Search(req filter.Request) ([]Tag, int, error)
	ReadAll() ([]Tag, error)
	GetTotalReposByTag(item Tag) (int, error)
    Remove(item Tag) error
    Update(item Tag) error
}

type RepoRepository interface {
	Create(item Repo) error
	ReadOne(item Repo) (Repo, error)
	Search(req filter.Request) ([]Repo, int, error)
	SearchByTagsIDs(tags []string, filters filter.Request) ([]Repo, int, error)
	GetTagByName(name string) (Tag, error)
	RemoveTagFromRepo(repo Repo, tag Tag) error
	AddTagToRepo(repo Repo, tag Tag) error
    GetTotal() int
}

type RepoService interface {
	Create(item Repo) error
	ReadOne(item Repo) (Repo, error)
	Search(req filter.Request) ([]Repo, int, error)
	SearchByTagsIDs(tags []string, filters filter.Request) ([]Repo, int, error)
	GetTotalStarredRepositories() (int64, error)
	Sync() ([]Repo, int64, error)
	RemoveTagFromRepo(repo Repo, tag Tag) error
	AddTagToRepo(repo Repo, tag Tag) (Tag, error)
}
