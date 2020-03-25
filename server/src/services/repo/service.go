package repo

import (
	"server/libs/configuration"
	"server/libs/filter"
	"server/libs/github"
	"server/libs/logger"
	"server/src/containers/repository"
	"server/src/domain"
)

type RepoService struct {
	repoRepository domain.RepoRepository
	githubClient   github.GithubClient
}

func NewRepoService(repositories *repository.RepositoryContainer) domain.RepoService {
	token := configuration.Load().GithubPersonalToken
	githubClient := github.NewGithubClient(token)
	return &RepoService{
		repoRepository: repositories.RepoRepository,
		githubClient:   githubClient,
	}
}

func (s *RepoService) Sync() ([]domain.Repo, error) {
	var repos []domain.Repo
	starredRepositories, err := s.githubClient.FetchAllStarredRepositories([]github.Repository{}, "")
	if err != nil {
		logger.Error("Cannot sync with GitHub:", err)
		return repos, err
	}
	logger.Info("Got", len(starredRepositories), "starred repositories.")
	for _, starred := range starredRepositories {
		repo := domain.Repo{
			ID:          starred.ID,
			Name:        starred.Name,
			Description: starred.Description,
			URL:         starred.URL,
			Homepage:    starred.Homepage,
		}
		_, err := s.ReadOne(repo)
		if err == nil {
			continue
		}
		if err := s.Create(repo); err != nil {
			logger.Error("Cannot save the repository", repo.Name, "error=", err)
			continue
		}
		repos = append(repos, repo)
	}
	return repos, err
}

func (s *RepoService) Create(item domain.Repo) error {
	return s.repoRepository.Create(item)
}

func (s *RepoService) Update(item domain.Repo) error {
	return s.repoRepository.Update(item)
}

func (s *RepoService) ReadOne(item domain.Repo) (domain.Repo, error) {
	return s.repoRepository.ReadOne(item)
}

func (s *RepoService) Search(item filter.Request) ([]domain.Repo, int, error) {
	return s.repoRepository.Search(item)
}
