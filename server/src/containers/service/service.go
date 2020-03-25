package service

import (
	"server/src/containers/repository"
	"server/src/domain"
	"server/src/services/repo"
)

type ServiceContainer struct {
	RepoService domain.RepoService
}

func NewServiceContainer(repositories *repository.RepositoryContainer) *ServiceContainer {
	return &ServiceContainer{
		RepoService: repo.NewRepoService(repositories),
	}
}
