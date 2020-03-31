package repository

import (
	"backend/src/domain"
	"backend/src/sqlite"
	"backend/src/sqlite/repo"
	"backend/src/sqlite/tag"
)

type RepositoryContainer struct {
	RepoRepository domain.RepoRepository
	TagRepository  domain.TagRepository
}

func NewRepositoryContainer(db *sqlite.Database) *RepositoryContainer {
	return &RepositoryContainer{
		RepoRepository: repo.NewRepoRepository(db),
		TagRepository:  tag.NewTagRepository(db),
	}
}
