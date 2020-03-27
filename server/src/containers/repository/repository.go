package repository

import (
	"server/src/domain"
	"server/src/sqlite"
	"server/src/sqlite/repo"
	"server/src/sqlite/tag"
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
