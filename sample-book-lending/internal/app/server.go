package app

import (
	"context"
	"fmt"

	//"log"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/protobuf/encoding/protojson"

	"sample-book-lending/internal/data"    // internal/dataのパスに置き換えてください。
	"sample-book-lending/internal/service" // internal/serviceのパスに置き換えてください。
	pb "sample-book-lending/pkg/grpc"      // protobuf生成されたコードのパスに置き換えてください。
)

type server struct {
	pb.UnimplementedLibraryServiceServer
	pb.UnimplementedAccountServiceServer
	libraryService *service.LibraryService
	accountService *service.AccountService
}

func NewServer() *server {
	bookDB := data.NewGoLevelDB("./books.db")
	accountDB := data.NewAccountDB("./accounts.db") // アカウント用のDBを初期化

	b := data.Book{Id: "123", Copy: 10}
	bookDB.AddBook(b)

	return &server{
		libraryService: service.NewLibraryService(bookDB),
		accountService: service.NewAccountService(accountDB), // アカウントサービスを初期化
	}
}

func (s *server) BorrowBook(ctx context.Context, req *pb.BorrowBookRequest) (*pb.BorrowBookResponse, error) {
	fmt.Println("debug: s.libraryService.BorrowBook(ctx, req) is called")
	fmt.Println(req)
	return s.libraryService.BorrowBook(ctx, req)
}

func (s *server) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	fmt.Println("debug: s.accountService.RegisterUser(ctx, req) is called")
	fmt.Println(req)
	return s.accountService.RegisterUser(ctx, req)
}

func (s *server) GetUserInfo(ctx context.Context, req *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	fmt.Println("debug: s.accountService.GetUserInfo(ctx, req) is called")
	fmt.Println(req)
	return s.accountService.GetUserInfo(ctx, req)
}

func (s *server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	fmt.Println("debug: s.accountService.GetUserInfo(ctx, req) is called")
	fmt.Println(req)
	return s.accountService.LoginUser(ctx, req)
}

// 他のRPCメソッドも同様に実装

func NewGatewayMux() *runtime.ServeMux {
	return runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	}))
}

func GetContext() context.Context {
	return context.Background()
}
