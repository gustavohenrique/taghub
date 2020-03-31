package tag

import (
	"testing"

	"backend/libs/testutils/assert"
	"backend/src/domain"
	"backend/src/sqlite"
	"backend/src/sqlite/test"
)

func TestCreateTag(ts *testing.T) {
	test.Seed(ts, "Should save all data", func(t *testing.T, db *sqlite.Database) {
		item := domain.Tag{
			Name: "linux",
		}
		saved, err := NewTagRepository(db).Create(item)
		assert.Nil(t, err, "")
		assert.True(t, saved.ID != "", saved.ID)
	})

	test.Seed(ts, "Should fail when duplicate name", func(t *testing.T, db *sqlite.Database) {
		item := domain.Tag{
			Name: "golang",
		}
		_, err := NewTagRepository(db).Create(item)
		assert.NotNil(t, err, "")
	})
}
