package service

import (
	"server/src/containers/repository"
	"server/src/domain"
	"server/src/services/repo"
	"server/src/services/tag"
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
