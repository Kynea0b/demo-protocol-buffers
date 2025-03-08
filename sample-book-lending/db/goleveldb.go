package db

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
	"log"
	"strconv"
)

type GoLevelDB struct {
	Db *leveldb.DB
}

func NewGoLevelDB(path string) *GoLevelDB {
	db, _ := leveldb.OpenFile(path, nil)
	return &GoLevelDB{Db: db}
}

type Book struct {
	Title string
	Copy  int
}

// var accountdb *leveldb.DB
// 本の追加
func (dbBook *GoLevelDB) AddBook(book Book) {
	err := dbBook.Db.Put([]byte(book.Title), []byte(strconv.Itoa(book.Copy)), nil)
	if err != nil {
		log.Printf("failed to put book: %v", err)
	}
}

func (dbBook *GoLevelDB) AddBooks(books []Book) {
	for _, b := range books {
		err := dbBook.Db.Put([]byte(b.Title), []byte(strconv.Itoa(b.Copy)), nil)
		if err != nil {
			log.Printf("failed to put book: %v", err)
		}
	}
}

func (dbBook *GoLevelDB) BorrowBook(title string) {
	copyBytes, err := dbBook.Db.Get([]byte(title), nil)
	if err != nil {
		log.Printf("book '%s' not found: %v", title, err)
		return
	}
	copies, err := strconv.Atoi(string(copyBytes))
	if copies > 0 {
		copies--
		err := dbBook.Db.Put([]byte(title), []byte(strconv.Itoa(copies)), nil)
		if err != nil {
			log.Printf("failed to put book: %v", err)
		}
		fmt.Printf("Book '%s' borrowed\n", title)
	} else {
		fmt.Printf("Book '%s' is out of stock\n", title)
	}
}

// 　本の数取得
func (dbBook *GoLevelDB) GetBookCopies(title string) (int, error) {
	copiesBytes, err := dbBook.Db.Get([]byte(title), nil)
	if err != nil {
		return 0, fmt.Errorf("book '%s' not found: %v", title, err)
	}

	copies, err := strconv.Atoi(string(copiesBytes))
	if err != nil {
		return 0, fmt.Errorf("failed to convert copies to int: %v", err)
	}

	return copies, nil
}

// 本の削除
func (dbBook *GoLevelDB) DeleteItem(key string) {
	_ = dbBook.Db.Delete([]byte(key), nil)
}

func (dbBook *GoLevelDB) UpdateBookLendingCard(title string, name string) {
	// todo: panic occurs when the key does not exist
	// タイトル前方一致で取得
	iter := dbBook.Db.NewIterator(util.BytesPrefix([]byte(title)), nil)
	var key []byte
	for iter.Next() {
		//
		value := iter.Value()
		if len(value) == 0 {
			fmt.Println("貸し出し可")
			key = iter.Key()
			break
		} else {
			fmt.Println("貸し出し中")
		}
	}

	// 貸す本
	fmt.Println("Lend this book: ", string(key))

	// 貸与者の名前を書き込み
	err := dbBook.Db.Put(key, []byte(name), nil)
	if err != nil {
		fmt.Println("DB Error")
		return
	}
}
