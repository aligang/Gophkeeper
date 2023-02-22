package fs

import (
	"context"
	"crypto/md5"
	"github.com/aligang/Gophkeeper/internal/common/logging"
	"github.com/aligang/Gophkeeper/internal/server/config"
	"os"
)

type FileRepository struct {
	root   string
	logger *logging.InternalLogger
}

func New(s *config.Config) *FileRepository {
	repo := &FileRepository{
		root:   s.FileStorage,
		logger: logging.Logger.GetSubLogger("FileRepository", "file"),
	}
	repo.logger.Debug("Initialization file storage")
	err := os.MkdirAll(repo.root, 0755)
	if err != nil {
		repo.logger.Fatal("File storage initialization failed :%s", err.Error())
	}
	repo.logger.Debug("File storage initialization succeeded")
	return repo
}

func (r *FileRepository) Save(ctx context.Context, objectName string, object []byte) error {
	logger := r.logger.GetSubLogger("Object", "Save")
	storageFilePath := r.root + "/" + objectName
	logger.Debug("saving object %s to file %s", objectName, storageFilePath)
	err := SaveFile(ctx, storageFilePath, object)
	if err != nil {
		logger.Crit("Filed to save file: %s", storageFilePath)
		return err
	}
	logger.Debug("file written: %s", storageFilePath)
	return nil
}

func (r *FileRepository) Read(ctx context.Context, objectName string) ([]byte, error) {
	logger := r.logger.GetSubLogger("Object", "Read")
	logger.Debug("Reading file %s", objectName)
	storageFilePath := r.root + "/" + objectName
	data, err := ReadFile(ctx, storageFilePath)
	if err != nil {
		logger.Crit("Failed")
		return nil, err
	}
	logger.Debug("Succeeded")
	return data, nil
}

func (r *FileRepository) Delete(ctx context.Context, objectName string) error {
	logger := r.logger.GetSubLogger("Object", "Delete")
	filePath := r.root + "/" + objectName
	logger.Debug("Deleting file from file repository %s", filePath)
	err := DeleteFile(ctx, filePath)
	if err != nil {
		logger.Crit("Could not delete file repository %s", filePath)
		return err
	}
	r.logger.Debug("File successfully deleted %s", filePath)
	return nil
}

func (r *FileRepository) ListHashes(ctx context.Context) (map[string]any, error) {
	logger := r.logger.GetSubLogger("Object", "List")
	logger.Debug("Listing content of file repository")
	files, err := os.ReadDir(r.root)
	repoContent := map[string]any{}
	if err != nil {
		logger.Fatal("Error during listing directory content: %s, %s", r.root, err.Error())
	}
	for _, f := range files {
		content, err := r.Read(ctx, f.Name())
		if err != nil {
			logger.Crit("Error reading object from filesystem")
			return repoContent, err
		}
		md5Hash := md5.Sum(content)
		repoContent[string(md5Hash[:])] = nil
	}

	return repoContent, nil
}

func (r *FileRepository) Equals(another *FileRepository) bool {
	hashes, err := r.ListHashes(context.Background())
	if err != nil {
		return false
	}
	anotherHashes, err := another.ListHashes(context.Background())
	if err != nil {
		return false
	}
	if len(hashes) != len(anotherHashes) {
		return false
	}
	for key, _ := range hashes {
		if _, ok := anotherHashes[key]; !ok {
			return false
		}
	}
	return true
}

func (r *FileRepository) GetRootDirectory() string {
	return r.root
}

func (r *FileRepository) CleanOut() error {
	return os.RemoveAll(r.root)
}
