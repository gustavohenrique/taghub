package service

import (
	"backend/src/containers/repository"
	"backend/src/domain"
	"backend/src/services/repo"
	"backend/src/services/tag"
)

type ServiceContainer struct {
	RepoService domain.RepoService
	TagService  domain.TagService
}

func NewServiceContainer(repositories *repository.RepositoryContainer) *ServiceContainer {
	return &ServiceContainer{
		RepoService: repo.NewRepoService(repositories),
		TagService:  tag.NewTagService(repositories),
	}
}
