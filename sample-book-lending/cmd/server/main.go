package main

import (
	mydb "sample-book-lending/db"
	myutil "sample-book-lending/util"

	"github.com/syndtr/goleveldb/leveldb/util"

	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	hellopb "sample-book-lending/pkg/grpc"

	"google.golang.org/grpc/reflection"

	// (一部抜粋)
	"net"

	"google.golang.org/grpc"

	// for proxy
	"net/http"
	"unsafe"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	//"github.com/syndtr/goleveldb/leveldb"
	"google.golang.org/grpc/credentials/insecure"

	// for timestamp
	tspb "google.golang.org/protobuf/types/known/timestamppb"
)

// key: "本のタイトル" + "本のid", val: "貸与者の名前"
// key: "貸与者の名前", val: "貸し出し日"
var dbBook = mydb.NewGoLevelDB("path/to/bookdb")

// TODO
//// key: "name", val: "id"
////　db package関数整理しないと。本だけになっている。
//var dbAccount = mydb.NewGoLevelDB("path/to/accountsdb")

func init() {

}

type myServer struct {
	hellopb.UnimplementedLendingBooksServiceServer
}

// time.Time型 -> 文字列
func time2byteArray(t *tspb.Timestamp) []byte {
	return []byte(t.AsTime().String())
}

// SendBorrow account.protoの`service`に定義したメソッドの実装
// 本を借りるためのメソッド
func (s *myServer) SendBorrow(ctx context.Context, req *hellopb.BorrowRequest) (*hellopb.BorrrowResponse, error) {
	// 貸し出し表を更新
	dbBook.UpdateBookLendingCard(req.Book.Title, req.Account.Name)

	time := tspb.Now()
	dbBook.Db.Put([]byte(req.Account.Name), time2byteArray(time), nil)

	return &hellopb.BorrrowResponse{
		Account:   &hellopb.Account{Name: req.Account.Name},
		Book:      &hellopb.Book{Title: req.Book.Title},
		Timestamp: time,
	}, nil
}

// RegisterBook 本のタイトルから貸与者を取得
func (s *myServer) RegisterBook(ctx context.Context, req *hellopb.RegisterBookRequest) (*hellopb.RegisterBookResponse, error) {
	dbBook.RegisterBook(req.Title, int(req.Num))

	return &hellopb.RegisterBookResponse{
		Num:   req.Num,
		Title: req.Title,
	}, nil
}

// GetLendingInfo 本のタイトルから貸与者を取得
func (s *myServer) GetLendingInfo(ctx context.Context, req *hellopb.Book) (*hellopb.Accounts, error) {
	iter := dbBook.Db.NewIterator(util.BytesPrefix([]byte(req.Title)), nil)
	var acntArray []*hellopb.Account
	for iter.Next() {
		if len(iter.Value()) != 0 {
			// names = append(names, string(iter.Value()))
			acntArray = append(acntArray, &hellopb.Account{Name: string(iter.Value())})
		}
	}

	return &hellopb.Accounts{
		Accounts: acntArray,
	}, nil
}

func (s *myServer) GetBorrowedTime(ctx context.Context, req *hellopb.Account) (*hellopb.BorrrowResponse, error) {
	data, _ := dbBook.Db.Get([]byte(req.Name), nil)
	time_str := *(*string)(unsafe.Pointer(&data))

	t := tspb.New(myutil.StringToTime(time_str))

	// todo
	// specify book name

	return &hellopb.BorrrowResponse{
		Account:   &hellopb.Account{Name: req.Name},
		Timestamp: t,
	}, nil
}

//	{
//	 "body": {
//	   "name": "Hoge",
//	   "mail": "hogehoge@gmail.com"
//	 },
//	 // ... other fields
//	}
func (s *myServer) RegisterAccount(ctx context.Context, req *hellopb.AccountRequest) (*hellopb.AccountResponse, error) {
	fmt.Printf("Hello!! My Hobby is %s", req.Hobby)
	fmt.Printf("Hello!! My Hobby is %s", req.Body)

	return &hellopb.AccountResponse{
		AccountId: "123456789",
	}, nil
}

// NewMyServer 自作サービス構造体のコンストラクタを定義
func NewMyServer() *myServer {
	return &myServer{}
}

func main() {

	// 貸し出し書籍各3冊
	b1 := mydb.Book{Title: "赤毛のアン", Num: 3}
	b2 := mydb.Book{Title: "小公女セーラ", Num: 3}
	b3 := mydb.Book{Title: "フランダースの犬", Num: 3}

	books := []mydb.Book{b1, b2, b3}

	// 本の初期登録
	dbBook.RegisterBooks(books)

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

	err = gwmux.HandlePath("GET", "/hello/{name}", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		w.Write([]byte("hello " + pathParams["name"]))
	})

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
