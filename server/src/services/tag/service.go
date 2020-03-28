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

func (s *TagService) Search(item filter.Request) ([]domain.Tag, int, error) {
	return s.tagRepository.Search(item)
}
