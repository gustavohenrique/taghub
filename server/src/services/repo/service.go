package repo

import (
	"server/libs/filter"
	"server/src/containers/repository"
	"server/src/domain"
)

type RepoService struct {
	RepoRepository domain.RepoRepository
}

func NewRepoService(repositories *repository.RepositoryContainer) domain.RepoService {
	return &RepoService{repositories.RepoRepository}
}

func (s *RepoService) Create(item domain.Repo) (domain.Repo, error) {
	return s.RepoRepository.Create(item)
}

func (s *RepoService) Update(item domain.Repo) (domain.Repo, error) {
	return s.RepoRepository.Update(item)
}

func (s *RepoService) ReadOne(item domain.Repo) (domain.Repo, error) {
	return s.RepoRepository.ReadOne(item)
}

func (s *RepoService) Search(item filter.Request) ([]domain.Repo, int, error) {
	return s.RepoRepository.Search(item)
}
