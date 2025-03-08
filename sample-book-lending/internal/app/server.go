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
	pb.UnimplementedLendingBooksServiceServer
	libraryService *service.LibraryService
}

func NewServer() *server {
	db := data.NewGoLevelDB("./books.db")
	return &server{
		libraryService: service.NewLibraryService(db),
	}
}

func (s *server) SendBorrow(ctx context.Context, req *pb.BorrowRequest) (*pb.BorrrowResponse, error) {
	fmt.Println("debug: s.libraryService.BorrowBook(ctx, req) is called")
	fmt.Println(req)
	return s.libraryService.BorrowBook(ctx, req)
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
