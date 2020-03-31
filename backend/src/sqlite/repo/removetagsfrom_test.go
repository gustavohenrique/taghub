package repo

import (
	"testing"

	"backend/libs/testutils/assert"
	"backend/src/domain"
	"backend/src/sqlite"
	"backend/src/sqlite/test"
)

func TestRemoveTagFromRepo(ts *testing.T) {
	test.Seed(ts, "Should remove tag using tagID", func(t *testing.T, db *sqlite.Database) {
		repo := domain.Repo{ID: "MDEwOlJlcG9zaXRvcnkzMDY1NDU0"}
		tag := domain.Tag{ID: "1"}
		err := NewRepoRepository(db).RemoveTagFromRepo(repo, tag)
		assert.Nil(t, err, "")
	})
}
