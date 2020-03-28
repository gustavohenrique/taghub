package tag

import (
	"server/libs/filter"
	"server/src/containers/repository"
	"server/src/domain"
)

type TagService struct {
	tagRepository domain.TagRepository
}

func NewTagService(repositories *repository.RepositoryContainer) domain.TagService {
	return &TagService{
		tagRepository: repositories.TagRepository,
	}
}

func (s *TagService) ReadAll() ([]domain.Tag, error) {
	return s.tagRepository.ReadAll()
}

func (s *TagService) Remove(item domain.Tag) error {
	return s.tagRepository.Remove(item)
}

func (s *TagService) Update(item domain.Tag) error {
	return s.tagRepository.Update(item)
}

func (s *TagService) GetTotalReposByTag(item domain.Tag) (int, error) {
	return s.tagRepository.GetTotalReposByTag(item)
}

func (s *TagService) Search(item filter.Request) ([]domain.Tag, int, error) {
	return s.tagRepository.Search(item)
}
