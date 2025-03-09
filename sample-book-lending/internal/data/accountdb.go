package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
)

type AccountDB struct {
	Store KeyValueStore
}

func NewAccountDB(store KeyValueStore) *AccountDB {
	return &AccountDB{Store: store}
}

type User struct {
	ID       string
	Username string
	Password []byte
	Email    string
}

func (db *AccountDB) AddUser(user *User) (string, error) {
	exists, err := db.UserExists(user.Username, user.Email)
	if err != nil {
		return "", err
	}

	if exists {
		return "", ErrAlreadyExists
	}

	id := uuid.New().String()
	user.ID = id

	userData, err := json.Marshal(user)
	if err != nil {
		return "", fmt.Errorf("failed to marshal user data: %w", err)
	}

	err = db.Store.Put([]byte(id), userData)
	if err != nil {
		return "", ErrInternal
	}
	return id, nil
}

func (db *AccountDB) GetUser(id string) (*User, error) {
	fmt.Println("GetUser called with id:", id)

	data, err := db.Store.Get([]byte(id))
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("failed to get user data from store: %w", err)
	}

	var user User
	err = json.Unmarshal(data, &user)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal user data: %w", err)
	}

	return &user, nil
}

func (db *AccountDB) UserExists(username, email string) (bool, error) {
	//LevelDBStoreインスタンスを取得
	levelDBStore, ok := db.Store.(*LevelDBStore)
	if !ok {
		return false, errors.New("store is not LevelDBStore")
	}
	//LevelDBのDBインスタンスを取得
	iter := levelDBStore.DB.NewIterator(nil, nil)
	defer iter.Release()
	for iter.Next() {
		// ログを追加
		log.Printf("Raw data from LevelDB: %s", iter.Value())

		var user User
		err := json.Unmarshal(iter.Value(), &user)
		if err != nil {
			log.Printf("failed to unmarshal user data: %v, raw data: %s", err, iter.Value())
			return false, fmt.Errorf("failed to unmarshal user data: %w", err)
		}

		if user.Username == username && user.Email == email {
			return true, nil
		}
	}
	return false, nil
}

func (db *AccountDB) GetUserByUsername(username string) (*User, error) {
	//LevelDBStoreインスタンスを取得
	levelDBStore, ok := db.Store.(*LevelDBStore)
	if !ok {
		return nil, errors.New("store is not LevelDBStore")
	}
	//LevelDBのDBインスタンスを取得
	iter := levelDBStore.DB.NewIterator(nil, nil)
	defer iter.Release()
	for iter.Next() {
		var user User
		err := json.Unmarshal(iter.Value(), &user)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal user data: %w", err)
		}

		if user.Username == username {
			return &user, nil
		}
	}
	return nil, ErrNotFound
}
