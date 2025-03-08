package service

import (
	"context"
	"fmt"
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

func (s *LibraryService) BorrowBook(ctx context.Context, req *pb.BorrowBookRequest) (*pb.BorrowBookResponse, error) {
	// ビジネスロジックを実装
	fmt.Println("service; ", req)
	err := s.db.DecrementBookCopies(req.BookId)
	if err != nil {
		return nil, fmt.Errorf("hogehogeerr")
	}

	// 現在時刻の Timestamp を生成
	ts := tspb.Now()

	// Timestamp を time.Time に変換
	t := ts.AsTime()

	// 年月日に変換して文字列として出力
	//dateStr := t.Format("2006-01-02")
	//時刻を秒単位まで入れる
	dateStr := t.Format("2006-01-02 15:04:05")

	return &pb.BorrowBookResponse{
		Message: dateStr,
	}, nil
}
