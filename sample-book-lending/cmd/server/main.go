package main

import (
	util "sample-book-lending/util"

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

func UpdateStock(key string, db *leveldb.DB) {
	// todo: panic occurs when the key does not exist
	data, err := db.Get([]byte(key), nil)
	if err != nil {
		fmt.Println("DB Error")
		return
	}
	val := util.DecodeUint(data)

	// 貸し出し数を減らす
	if val == 0 {
		fmt.Println("No Stock")
		return
	}
	val = val - 1

	// 残りの本の数
	fmt.Println("Stock remaining: ", val)

	// update
	buf, _ := util.EncodeUint(uint64(val))
	err = db.Put([]byte(key), buf, nil)
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

	// 本の数を1冊減らす
	UpdateStock(req.Book.Title, dbBook)

	return &hellopb.BorrrowResponse{
		Account: &hellopb.Account{Name: req.Account.Name},
		Book:    &hellopb.Book{Title: req.Book.Title},
	}, nil
}

func (s *myServer) ShowAccountInfo(ctx context.Context, req *hellopb.Account) (*hellopb.AccountInfo, error) {
	var titles []string
	titles = append(titles, "foo title")
	titles = append(titles, "bar title")
	books := &hellopb.Books{Titles: titles}
	return &hellopb.AccountInfo{
		Name:  req.GetName(),
		Books: books,
	}, nil
}

// 自作サービス構造体のコンストラクタを定義
func NewMyServer() *myServer {
	return &myServer{}
}

func main() {
	// bool library
	dbBook, _ = leveldb.OpenFile("path/to/bookdb", nil)
	// 書き込み
	// 貸し出し書籍10冊
	buf, _ := util.EncodeUint(uint64(10))
	_ = dbBook.Put([]byte("赤毛のアン"), buf, nil)
	_ = dbBook.Put([]byte("小公女セーラ"), buf, nil)
	_ = dbBook.Put([]byte("フランダースの犬"), buf, nil)

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
