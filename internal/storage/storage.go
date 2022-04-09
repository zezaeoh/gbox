package storage

import (
	"errors"
	"fmt"
	"github.com/shivamMg/ppds/tree"
	"github.com/zezaeoh/gbox/internal/storage/github"
)

type Storage interface {
	Get(name string) (string, error)
	Set(name, data string) error
	Delete(name string) error
	List() (tree.Node, error)
	GetMatched(toMatched string) []string
}

var (
	// current storage
	_storage   Storage
	storageMap map[string]Storage
)

// exported methods
var (
	GetStorageMap = getStorageMap
	GetStorage    = getStorage
)

func getStorageWithName(name string) (Storage, error) {
	sm, err := getStorageMap()
	if err != nil {
		return nil, err
	}
	s, ok := sm[name]
	if !ok {
		return nil, fmt.Errorf("there is no storage named %s", name)
	}
	return s, nil
}

func getStorageMap() (map[string]Storage, error) {
	if storageMap != nil {
		return storageMap, nil
	}
	cfg, err := loadConfig()
	if err != nil {
		return nil, err
	}
	if len(cfg.Storages) == 0 {
		return nil, errors.New("there is no storage configs")
	}
	sm := make(map[string]Storage, len(cfg.Storages))
	for _, s := range cfg.Storages {
		switch s.Kind {
		case "github":
			sm[s.Name], err = github.NewStorage(s.Spec)
			if err != nil {
				return nil, err
			}
		}
	}
	storageMap = sm
	return sm, nil
}

func getStorage() (Storage, error) {
	if _storage != nil {
		return _storage, nil
	}
	cfg, err := loadConfig()
	if err != nil {
		return nil, err
	}
	name := cfg.CurrentStorage
	if len(name) == 0 {
		for _, s := range cfg.Storages {
			name = s.Name
			break
		}
	}
	s, err := getStorageWithName(name)
	if err != nil {
		return nil, err
	}
	_storage = s
	return s, nil
}
