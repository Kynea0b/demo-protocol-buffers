package data

import (
	"github.com/syndtr/goleveldb/leveldb"
)

type LevelDBStore struct {
	DB *leveldb.DB
}

func NewLevelDBStore(path string) (*LevelDBStore, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}
	return &LevelDBStore{DB: db}, nil
}

func (s *LevelDBStore) Get(key []byte) ([]byte, error) {
	data, err := s.DB.Get(key, nil)
	if err != nil {
		if err == LevelDBNotFound { // LevelDBNotFound を使用
			return nil, ErrNotFound
		}
		return nil, err
	}
	return data, nil
}

func (s *LevelDBStore) Put(key, value []byte) error {
	return s.DB.Put(key, value, nil)
}

func (s *LevelDBStore) Has(key []byte) (bool, error) {
	_, err := s.DB.Get(key, nil)
	if err != nil {
		if err == LevelDBNotFound { // LevelDBNotFound を使用
			return false, nil
		}
		return false, err
	}
	return true, nil
}
