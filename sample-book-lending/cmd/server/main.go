package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/reflection"
	hellopb "sample-book-lending/pkg/grpc"
	"log"
	"os"
	"os/signal"

	// (一部抜粋)
	"google.golang.org/grpc"
	"net"

	// for proxy
	"google.golang.org/grpc/credentials/insecure"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
)

type myServer struct {
	hellopb.UnimplementedLendingBooksServiceServer
}

// account.protoの`service`に定義したメソッドの実装
// 本を借りるためのメソッド
func (s *myServer) SendBorrow(ctx context.Context, req *hellopb.BorrowRequest) (*hellopb.BorrrowResponse, error) {
	// リクエストからnameフィールドを取り出して
	// "Hello, [名前]!"というレスポンスを返す
	return &hellopb.BorrrowResponse{
		Account: &hellopb.Account{Name: req.Account.Name},
		Book: &hellopb.Book{Title: req.Book.Title},
	}, nil
}

// account.protoの`service`に定義したメソッドの実装
// アカウントの貸与状態を知るためのRestful API
func (s *myServer) AccountInfo(ctx context.Context, req *hellopb.Account) (*hellopb.Account, error) {
	// リクエストからnameフィールドを取り出して
	// "Hello, [名前]!"というレスポンスを返す
	return &hellopb.Account{
		Name: req.GetName(),
	}, nil
}

// 自作サービス構造体のコンストラクタを定義
func NewMyServer() *myServer {
	return &myServer{}
}

func main() {
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
