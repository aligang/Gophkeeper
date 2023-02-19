package fs

import (
	"context"
	"github.com/aligang/Gophkeeper/pkg/config"
	"github.com/aligang/Gophkeeper/pkg/fixtures"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const referenceObjectName = "test-object"
const referencePayload = "test-data"

func TestRepoMethods(t *testing.T) {
	repoPath := fixtures.PathFixture()
	fileRepo := New(
		&config.ServerConfig{
			FileStorage: repoPath,
		},
	)

	t.Run("SAVE object method test", func(t *testing.T) {
		err := fileRepo.Save(context.Background(), referenceObjectName, []byte(referencePayload))
		assert.Equal(t, nil, err)
		dirContent, err := os.ReadDir(repoPath)
		assert.Equal(t, nil, err)
		assert.Equal(t, len(dirContent), 1)
		data, err := os.ReadFile(repoPath + "/" + referenceObjectName)
		assert.Equal(t, nil, err)
		assert.Equal(t, referencePayload, string(data))
	})

	t.Run("READ object method test", func(t *testing.T) {
		data, err := fileRepo.Read(context.Background(), referenceObjectName)
		assert.Equal(t, nil, err)
		assert.Equal(t, referencePayload, string(data))
	})

	t.Run("DELETE object method test", func(t *testing.T) {
		anotherObjectName := "another-object"
		anotherObjectPayload := "another-payload"
		err := fileRepo.Save(context.Background(), anotherObjectName, []byte(anotherObjectPayload))
		assert.Equal(t, nil, err)
		err = fileRepo.Delete(context.Background(), referenceObjectName)
		assert.Equal(t, nil, err)
		dirContent, err := os.ReadDir(repoPath)
		assert.Equal(t, nil, err)
		assert.Equal(t, len(dirContent), 1)
		fileName := dirContent[0].Name()
		assert.Equal(t, anotherObjectName, fileName)
		data, _ := os.ReadFile(repoPath + "/" + anotherObjectName)
		assert.Equal(t, anotherObjectPayload, string(data))
	})

	_ = os.RemoveAll(repoPath)

}
