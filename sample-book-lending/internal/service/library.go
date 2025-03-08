package service

import (
	"context"
	tspb "google.golang.org/protobuf/types/known/timestamppb"

	"sample-book-lending/internal/data" // internal/dataのパスに置き換えてください。
	pb "sample-book-lending/pkg/grpc"   // protobuf生成されたコードのパスに置き換えてください。
)

type LibraryService struct {
	db *data.GoLevelDB
}

func NewLibraryService(db *data.GoLevelDB) *LibraryService {
	return &LibraryService{db: db}
}

func (s *LibraryService) BorrowBook(ctx context.Context, req *pb.BorrowRequest) (*pb.BorrrowResponse, error) {
	// ビジネスロジックを実装
	book := data.Book{Title: req.Book.Title} // 仮の例
	s.db.AddBook(book)
	time := tspb.Now()

	return &pb.BorrrowResponse{
		Account:   &pb.Account{Name: req.Account.Name},
		Book:      &pb.Book{Title: req.Book.Title},
		Timestamp: time,
	}, nil
}
