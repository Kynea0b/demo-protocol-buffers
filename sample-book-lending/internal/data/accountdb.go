package data

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

type AccountDB struct {
	Db *leveldb.DB
}

func NewAccountDB(path string) *AccountDB {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		log.Fatalf("failed to open accountdb: %v", err)
	}
	return &AccountDB{Db: db}
}

type User struct {
	ID       string
	Username string
	Password []byte
	Email    string
}

func (db *AccountDB) AddUser(user *User) (string, error) {
	id := uuid.New().String()
	user.ID = id
	err := db.Db.Put([]byte(id), []byte(user.Username+":"+string(user.Password)+":"+user.Email), nil)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (db *AccountDB) GetUser(id string) (*User, error) {
	fmt.Println("GetUser called with id:", id) // デバッグ用
	data, err := db.Db.Get([]byte(id), nil)
	if err != nil {
		if err == leveldb.ErrNotFound {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %v", err)
	}

	userData := string(data)
	parts := strings.Split(userData, ":")
	if len(parts) != 3 {
		return nil, errors.New("invalid user data format")
	}

	username := parts[0]
	password := parts[1]
	email := parts[2]

	return &User{
		ID:       id,
		Username: username,
		Password: []byte(password),
		Email:    email,
	}, nil
}

func (db *AccountDB) GetUserByUsername(username string) (*User, error) {
	fmt.Println("GetUserByUsername called with username:", username) // デバッグ用

	iter := db.Db.NewIterator(util.BytesPrefix([]byte("")), nil)
	defer iter.Release() // deferでiter.Release()を確実に実行

	for iter.Next() {
		userData := string(iter.Value())
		parts := strings.Split(userData, ":")

		if len(parts) != 3 {
			fmt.Println("Invalid user data format:", userData) // デバッグ用
			continue                                           // 次のレコードへ
		}

		u := parts[0]
		p := parts[1]
		e := parts[2]

		if u == username {
			id := string(iter.Key())
			fmt.Println("User found:", username, "ID:", id) // デバッグ用

			return &User{
				ID:       id,
				Username: u,
				Password: []byte(p),
				Email:    e,
			}, nil
		}
	}

	fmt.Println("User not found:", username) // デバッグ用
	return nil, errors.New("user not found")
}
