package fs

import (
	"github.com/aligang/Gophkeeper/internal/config"
	"github.com/aligang/Gophkeeper/internal/fixtures"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestRepoInitializationGeneration(t *testing.T) {
	validExistingPath := "/tmp/" + fixtures.PathFixture()
	validNonExistingPath := "/tmp/" + fixtures.PathFixture()

	invalidPath := "/" + fixtures.PathFixture()
	_ = os.Remove(validExistingPath)
	_ = os.Remove(validNonExistingPath)
	_ = os.Remove(invalidPath)
	os.MkdirAll(validExistingPath, 0755)

	t.Run("succeed with valid existing path", func(t *testing.T) {

		assert.NotPanics(t, func() {
			New(
				&config.ServerConfig{
					FileStorage: validExistingPath,
				},
			)
		})
	})
	_ = os.RemoveAll(validExistingPath)

	t.Run("succeed with valid non-existing path", func(t *testing.T) {

		assert.NotPanics(t, func() {
			New(
				&config.ServerConfig{
					FileStorage: validExistingPath,
				},
			)
		})
	})
	_ = os.RemoveAll(validExistingPath)

	t.Run("failed with invalid path", func(t *testing.T) {

		assert.Panics(t, func() {
			New(
				&config.ServerConfig{
					FileStorage: invalidPath,
				},
			)
		})
	})

}
