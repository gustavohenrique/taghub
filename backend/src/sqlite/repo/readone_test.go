package repo

import (
	"testing"

	"backend/libs/testutils/assert"
	"backend/src/domain"
	"backend/src/sqlite"
	"backend/src/sqlite/test"
)

func TestReadOneRepo(ts *testing.T) {
	test.Seed(ts, "Should find by ID", func(t *testing.T, db *sqlite.Database) {
		item := domain.Repo{
			ID: "MDEwOlJlcG9zaXRvcnkzMDY1NDU0",
		}
		found, err := NewRepoRepository(db).ReadOne(item)
		assert.Nil(t, err, "")
		assert.Equal(t, found.ID, item.ID, "")
		assert.Equal(t, found.Name, "impress.js", "")
		assert.Equal(t, found.Description, "Its a presentation framework", "")
		assert.Equal(t, found.URL, "https://github.com/impress/impress.js", "")
		assert.Equal(t, found.Homepage, "http://impress.js.org", "")
	})
}
