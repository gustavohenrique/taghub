package repository

import (
	"server/src/domain"
	"server/src/sqlite"
	"server/src/sqlite/repo"
)

type RepositoryContainer struct {
	RepoRepository      domain.RepoRepository
}

func NewRepositoryContainer(db *sqlite.Database) *RepositoryContainer {
	return &RepositoryContainer{
		RepoRepository:      repo.NewRepoRepository(db),
	}
}
