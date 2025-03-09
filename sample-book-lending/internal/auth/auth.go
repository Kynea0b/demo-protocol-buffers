package auth

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc/metadata"
)

// JWTのシークレットキー
var jwtSecretKey = []byte("your-secret-key")

// 認証トークンを検証するミドルウェア
func AuthMiddleware(ctx context.Context) (context.Context, error) {
	// メタデータから認証トークンを取得
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("metadata is not provided")
	}

	//authHeader, ok := md["authorization"]
	//if !ok || len(authHeader) == 0 {
	//	return nil, errors.New("authorization header is not provided")
	//}
	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "authorization header is not provided")
	}

	tokenString := strings.Replace(authHeader[0], "Bearer ", "", 1)

	// JWTトークンを検証
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})
	if err != nil {
		log.Printf("failed to parse token: %v", err)
		return nil, status.Errorf(codes.Unauthenticated, "invalid token")
	}

	// トークンが有効か確認
	if !token.Valid {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token")
	}

	// トークンからユーザーIDを取得
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token claims")
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "invalid user ID in token")
	}

	// コンテキストにユーザーIDを追加
	ctx = context.WithValue(ctx, "user_id", userID)

	return ctx, nil
}
