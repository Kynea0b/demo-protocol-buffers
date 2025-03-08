package service

import (
	"context"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"

	"sample-book-lending/internal/data"
	pb "sample-book-lending/pkg/grpc"
)

type AccountService struct {
	db *data.AccountDB
}

func NewAccountService(db *data.AccountDB) *AccountService {
	return &AccountService{db: db}
}

func (s *AccountService) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	// パスワードをハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("failed to hash password: %v", err)
		return nil, errors.New("failed to register user")
	}

	user := &data.User{
		Username: req.Username,
		Password: hashedPassword,
		Email:    req.Email,
	}

	// ユーザーをデータベースに登録
	userID, err := s.db.AddUser(user)
	if err != nil {
		log.Printf("failed to add user: %v", err)
		return nil, errors.New("failed to register user")
	}

	return &pb.RegisterUserResponse{UserId: userID}, nil
}

func (s *AccountService) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	// ユーザーをデータベースから取得
	user, err := s.db.GetUserByUsername(req.Username)
	if err != nil {
		log.Printf("failed to get user: %v", err)
		return nil, errors.New("failed to login")
	}

	// パスワードを検証
	err = bcrypt.CompareHashAndPassword(user.Password, []byte(req.Password))
	if err != nil {
		log.Printf("invalid password: %v", err)
		return nil, errors.New("failed to login")
	}

	// トークンを生成 (JWTなど)
	token, err := generateToken(user.Username)
	if err != nil {
		log.Printf("failed to generate token: %v", err)
		return nil, errors.New("failed to login")
	}

	return &pb.LoginUserResponse{UserId: user.ID, Token: token}, nil
}

func (s *AccountService) GetUserInfo(ctx context.Context, req *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	// ユーザーをデータベースから取得
	user, err := s.db.GetUser(req.UserId)
	if err != nil {
		log.Printf("failed to get user: %v", err)
		return nil, errors.New("failed to get user info")
	}

	return &pb.GetUserInfoResponse{Username: user.Username, Email: user.Email}, nil
}

// トークン生成関数 (JWTなど)
func generateToken(username string) (string, error) {
	// JWTトークン生成処理を実装
	return "dummy_token", nil
}
