package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc/reflection"

	"sample-book-lending/internal/auth"

	"google.golang.org/grpc"

	"sample-book-lending/internal/app"
	pb "sample-book-lending/pkg/grpc"
)

func main() {
	port_server := ":35635"
	// gRPCサーバーの起動
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(authInterceptor), // インターセプターを登録 auth追加のため
	)
	// リフレクションAPIの登録 (追加)
	reflection.Register(grpcServer)
	// サービスの登録
	server := app.NewServer()
	pb.RegisterLibraryServiceServer(grpcServer, server)
	pb.RegisterAccountServiceServer(grpcServer, server)
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve gRPC: %v", err)
		}
	}()

	// gRPC Gatewayの起動
	mux := app.NewGatewayMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	// rest api呼び出しに必要
	if err := pb.RegisterLibraryServiceHandlerFromEndpoint(app.GetContext(), mux, "localhost:50051", opts); err != nil {
		log.Fatalf("failed to register gateway: %v", err)
	}

	// rest api呼び出しに必要
	if err := pb.RegisterAccountServiceHandlerFromEndpoint(app.GetContext(), mux, "localhost:50051", opts); err != nil {
		log.Fatalf("failed to register gateway: %v", err)
	}

	// HTTPサーバーの起動
	httpServer := &http.Server{
		Addr:    port_server,
		Handler: mux,
	}
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalf("failed to serve HTTP: %v", err)
	}
}

// 認証インターセプター
func authInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// 認証が不要なメソッドはスキップ
	if info.FullMethod == "/account.AccountService/RegisterUser" || info.FullMethod == "/account.AccountService/LoginUser" {
		return handler(ctx, req)
	}

	// 認証トークンを検証
	newCtx, err := auth.AuthMiddleware(ctx)
	if err != nil {
		return nil, err
	}

	// コンテキストを更新して次のハンドラーを呼び出す
	return handler(newCtx, req)
}
