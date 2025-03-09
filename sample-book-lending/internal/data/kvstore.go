package data

import (
	"errors"

	"github.com/syndtr/goleveldb/leveldb"
)

var (
	ErrNotFound      = errors.New("not found")
	ErrAlreadyExists = errors.New("already exists")
	ErrInternal      = errors.New("internal error")
	LevelDBNotFound  = leveldb.ErrNotFound // leveldb.ErrNotFound を定義
)

type KeyValueStore interface {
	Get(key []byte) ([]byte, error)
	Put(key, value []byte) error
	Has(key []byte) (bool, error)
	// 必要に応じて他のメソッドを追加
}
