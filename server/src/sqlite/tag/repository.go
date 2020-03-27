package tag

import (
	"server/libs/errors"
	"server/src/domain"
	"server/src/sqlite"
)

type TagRepository struct {
	db *sqlite.Database
}

func NewTagRepository(db *sqlite.Database) domain.TagRepository {
	return TagRepository{db}
}

func (tag TagRepository) Create(item domain.Tag) (domain.Tag, error) {
	query := "INSERT INTO tags (name) VALUES (?)"
	id, err := tag.db.ExecAndGetLastID(query,
		item.Name,
	)
	if err != nil {
		code := errors.Detect(err)
		return item, errors.New(code, err.Error())
	}
	item.ID = id
	return item, nil
}
