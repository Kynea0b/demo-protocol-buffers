package db

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
	myutil "sample-book-lending/util"
	"unsafe"
)

type GoLevelDB struct {
	Db *leveldb.DB
}

func NewGoLevelDB(path string) *GoLevelDB {
	db, _ := leveldb.OpenFile(path, nil)
	return &GoLevelDB{Db: db}
}

// var accountdb *leveldb.DB
// 本の追加
func (dbBook *GoLevelDB) AddItem(key string, val string) {
	_ = dbBook.Db.Put([]byte(key), []byte(val), nil)
}

// 本の削除
func (dbBook *GoLevelDB) DeleteItem(key string) {
	_ = dbBook.Db.Delete([]byte(key), nil)
}

// 本の冊数取得
func (dbBook *GoLevelDB) GetItem(key string) string {
	data, _ := dbBook.Db.Get([]byte(key), nil)
	res := *(*string)(unsafe.Pointer(&data))
	return res
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

type Book struct {
	Title string
	Num   int
}

// 本のタイトルと冊数を指定してDB登録
func (dbBook *GoLevelDB) RegisterBook(title string, cnt int) {
	for i := 0; i < cnt; i++ {
		storekey := myutil.ParseStoreKey(title, i)
		// valueにはアカウントの`name`を登録
		// 初期登録では誰も借りていないので、空文字
		_ = dbBook.Db.Put(storekey, []byte(""), nil)
	}
}

func (dbBook *GoLevelDB) RegisterBooks(books []Book) {
	for _, b := range books {
		dbBook.RegisterBook(b.Title, b.Num)
	}
}
