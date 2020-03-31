package repo

import (
	"backend/libs/configuration"
	"backend/libs/filter"
	"backend/libs/github"
	"backend/libs/logger"
	"backend/src/containers/repository"
	"backend/src/domain"
)

type RepoService struct {
	repoRepository domain.RepoRepository
	tagRepository  domain.TagRepository
	githubClient   github.GithubClient
}

func NewRepoService(repositories *repository.RepositoryContainer) domain.RepoService {
	token := configuration.Load().GithubPersonalToken
	githubClient := github.NewGithubClient(token)
	return &RepoService{
		repoRepository: repositories.RepoRepository,
		tagRepository:  repositories.TagRepository,
		githubClient:   githubClient,
	}
}

func (s *RepoService) GetTotalStarredRepositories() (int64, error) {
	return s.githubClient.GetTotalStarredRepositories()
}

func (s *RepoService) Sync() ([]domain.Repo, int64, error) {
	var repos []domain.Repo
	starredRepositories, total, err := s.githubClient.FetchAllStarredRepositories([]github.Repository{}, "")
	if err != nil {
		logger.Error("Cannot sync with GitHub:", err)
		return repos, 0, err
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
	return repos, total, err
}

func (s *RepoService) Create(item domain.Repo) error {
	return s.repoRepository.Create(item)
}

func (s *RepoService) ReadOne(item domain.Repo) (domain.Repo, error) {
	return s.repoRepository.ReadOne(item)
}

func (s *RepoService) Remove(repo domain.Repo) error {
	return s.repoRepository.Remove(repo)
}

func (s *RepoService) SearchByTagsIDs(tags []string, item filter.Request) ([]domain.Repo, int, error) {
	return s.repoRepository.SearchByTagsIDs(tags, item)
}

func (s *RepoService) Search(item filter.Request) ([]domain.Repo, int, error) {
	return s.repoRepository.Search(item)
}

func (s *RepoService) AddTagToRepo(repo domain.Repo, tag domain.Tag) (domain.Tag, error) {
	if tag.ID == "" {
		item, err := s.tagRepository.Create(tag)
		if err != nil {
			logger.Info("Cannot create tag '", tag.Name, "'", err, ". Finding it...")
			item, err = s.tagRepository.ReadOne(tag)
			if err != nil {
				return tag, err
			}
		}
		tag = item
	}
	return tag, s.repoRepository.AddTagToRepo(repo, tag)
}

func (s *RepoService) RemoveTagFromRepo(repo domain.Repo, tag domain.Tag) error {
	return s.repoRepository.RemoveTagFromRepo(repo, tag)
}
