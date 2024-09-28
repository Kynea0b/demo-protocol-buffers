package main

import (
	"github.com/syndtr/goleveldb/leveldb/util"

	"context"
	"fmt"
	"google.golang.org/grpc/reflection"
	"log"
	"os"
	"os/signal"
	hellopb "sample-book-lending/pkg/grpc"

	// (一部抜粋)
	"google.golang.org/grpc"
	"net"

	// for proxy
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/syndtr/goleveldb/leveldb"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"unsafe"
)

var dbBook *leveldb.DB

const (
	title_book1 = "赤毛のアン"
	title_book2 = "小公女セーラ"
	title_book3 = "フランダースの犬"
)

var (
	num_book1 int
	num_book2 int
	num_book3 int
)

// var accountdb *leveldb.DB
// 本の追加
func addItem(key string, val string) {
	_ = dbBook.Put([]byte(key), []byte(val), nil)
}

// 本の削除
func deleteItem(key string) {
	_ = dbBook.Delete([]byte(key), nil)
}

// 本の冊数取得
func getItem(key string) string {
	data, _ := dbBook.Get([]byte(key), nil)
	res := *(*string)(unsafe.Pointer(&data))
	return res
}

func UpdateBookLendingCard(title string, name string, db *leveldb.DB) {
	// todo: panic occurs when the key does not exist
	// タイトル前方一致で取得
	iter := db.NewIterator(util.BytesPrefix([]byte(title)), nil)
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
	err := db.Put(key, []byte(name), nil)
	if err != nil {
		fmt.Println("DB Error")
		return
	}
}

type myServer struct {
	hellopb.UnimplementedLendingBooksServiceServer
}

// account.protoの`service`に定義したメソッドの実装
// 本を借りるためのメソッド
func (s *myServer) SendBorrow(ctx context.Context, req *hellopb.BorrowRequest) (*hellopb.BorrrowResponse, error) {
	// 貸し出し表を更新
	UpdateBookLendingCard(req.Book.Title, req.Account.Name, dbBook)

	return &hellopb.BorrrowResponse{
		Account: &hellopb.Account{Name: req.Account.Name},
		Book:    &hellopb.Book{Title: req.Book.Title},
	}, nil
}

// アカウントの名前から借りている本の一覧を取得
func (s *myServer) ShowBookInfo(ctx context.Context, req *hellopb.Book) (*hellopb.Accounts, error) {
	iter := dbBook.NewIterator(util.BytesPrefix([]byte(req.Title)), nil)
	var names []string
	for iter.Next() {
		if len(iter.Value()) != 0 {
			names = append(names, string(iter.Value()))
		}
	}
	return &hellopb.Accounts{
		Names: names,
	}, nil
}

// 自作サービス構造体のコンストラクタを定義
func NewMyServer() *myServer {
	return &myServer{}
}

// 本のタイトルからkeyに変換します
func parseStoreKey(key string, id int) []byte {
	storekey := fmt.Sprintf("%s:%d", key, id)
	return []byte(storekey)
}

// 本のタイトルと冊数を指定してDB登録
func registerBooks(title string, cnt int, dbBook *leveldb.DB) {
	for i := 0; i < cnt; i++ {
		storekey := parseStoreKey(title, i)
		// valueにはアカウントの`name`を登録
		// 初期登録では誰も借りていないので、空文字
		_ = dbBook.Put(storekey, []byte(""), nil)
	}
}

func main() {
	num_book1 = 3
	num_book2 = 3
	num_book3 = 3
	// bool library
	dbBook, _ = leveldb.OpenFile("path/to/bookdb", nil)
	// 書き込み
	// 貸し出し書籍各3冊
	registerBooks(title_book1, num_book1, dbBook)
	registerBooks(title_book2, num_book2, dbBook)
	registerBooks(title_book3, num_book3, dbBook)

	// 1. 8080番portのLisnterを作成
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	// 2. gRPCサーバーを作成
	s := grpc.NewServer()

	// 3. gRPCサーバーにGreetingServiceを登録
	// 第二引数はinterfaceであるGreetingServiceServerのため、これのメソッドリストを実装した構造体がはいる。
	hellopb.RegisterLendingBooksServiceServer(s, NewMyServer())

	// x_numberはproxy serverのインスタンス作成と起動です。
	// x_1. for proxy
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.NewClient(
		"0.0.0.0:8080",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	gwmux := runtime.NewServeMux()

	err = hellopb.RegisterLendingBooksServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	// 4. サーバーリフレクションの設定
	reflection.Register(s)

	// 5. 作成したgRPCサーバーを、8080番ポートで稼働させる
	go func() {
		log.Printf("start gRPC server port: %v", port)
		s.Serve(listener)
	}()

	// x_2. for proxy
	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())

	// 6.Ctrl+Cが入力されたらGraceful shutdownされるようにする
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}
