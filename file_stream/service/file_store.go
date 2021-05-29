package service

import (
	"bytes"
	"fmt"
	"os"
	"sync"

	"github.com/google/uuid"
)

//FileStore interface to store file
type FileStore interface {
	// save saves a new file
	Save(fileId string, fileType string, fileData bytes.Buffer) (string, error)
}

type DiskFileStore struct {
	mutex      sync.RWMutex
	fileFolder string
	files      map[string]*FileInfo
}

type FileInfo struct {
	FileID string
	Type   string
	Path   string
}

// returns a new DiskFileStore
func NewDiskFileStore(fileFolder string) *DiskFileStore {
	return &DiskFileStore{
		fileFolder: fileFolder,
		files:      make(map[string]*FileInfo),
	}
}

//  Save saves a new file to the store
func (store *DiskFileStore) Save(fileId string, fileType string, fileData bytes.Buffer) (string, error) {

	hashId, err := uuid.NewRandom()
	fileId = hashId.String()

	if err != nil {
		return "", fmt.Errorf("couldn't generate file id: %v", err)
	}
	filepath := fmt.Sprintf("%s/%s%s", store.fileFolder, fileId, fileType)
	file, err := os.Create(filepath)
	if err != nil {
		return "", fmt.Errorf("couldn't create file %v: ", err)
	}

	_, err = fileData.WriteTo(file)
	fmt.Println("hey man")
	fmt.Println(fileData.WriteTo(file))
	if err != nil {
		return "", fmt.Errorf("couldn't write file %v: ", err)
	}

	store.mutex.Lock()
	defer store.mutex.Unlock()

	store.files[fileId] = &FileInfo{
		FileID: fileId,
		Type:   fileType,
		Path:   filepath,
	}

	return fileId, nil
}
