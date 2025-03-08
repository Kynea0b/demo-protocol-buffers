package data

import (
	"fmt"
	"log"
	"strconv"

	"github.com/syndtr/goleveldb/leveldb"
)

type GoLevelDB struct {
	Db *leveldb.DB
}

func NewGoLevelDB(path string) *GoLevelDB {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		log.Fatalf("failed to open leveldb: %v", err)
	}
	return &GoLevelDB{Db: db}
}

func (dbBook *GoLevelDB) Close() {
	dbBook.Db.Close()
}

type Book struct {
	Id   string
	Copy int
}

// 本の追加
func (dbBook *GoLevelDB) AddBook(book Book) error {
	err := dbBook.Db.Put([]byte(book.Id), []byte(strconv.Itoa(book.Copy)), nil)
	if err != nil {
		return fmt.Errorf("failed to put book: %v", err)
	}
	return nil
}

// 本の取得
func (dbBook *GoLevelDB) GetBook(title string) (*Book, error) {
	copiesBytes, err := dbBook.Db.Get([]byte(title), nil)
	if err != nil {
		if err == leveldb.ErrNotFound {
			return nil, nil // 本が見つからない場合はnilを返す
		}
		return nil, fmt.Errorf("failed to get book: %v", err)
	}

	copies, err := strconv.Atoi(string(copiesBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to convert copies to int: %v", err)
	}

	return &Book{Id: title, Copy: copies}, nil
}

// 本の削除
func (dbBook *GoLevelDB) DeleteBook(title string) error {
	err := dbBook.Db.Delete([]byte(title), nil)
	if err != nil {
		return fmt.Errorf("failed to delete book: %v", err)
	}
	return nil
}

// 本の更新
func (dbBook *GoLevelDB) UpdateBook(book Book) error {
	return dbBook.AddBook(book) // AddBookと同じ処理
}

// 本の冊数を減らす
func (dbBook *GoLevelDB) DecrementBookCopies(id string) error {
	fmt.Println("dbbbbbbbbbbbbbbbb")
	book, err := dbBook.GetBook(id)
	fmt.Println("bookkkkkkk", book)
	if err != nil {
		return err
	}
	if book == nil {
		return fmt.Errorf("book '%s' not found", id)
	}

	if book.Copy > 0 {
		book.Copy--
		return dbBook.UpdateBook(*book)
	}
	return fmt.Errorf("book '%s' is out of stock", id)
}
