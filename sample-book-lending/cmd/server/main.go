package main

import (
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"

	"sample-book-lending/internal/app" // internal/appのパスに置き換えてください。
	pb "sample-book-lending/pkg/grpc"  // protobuf生成されたコードのパスに置き換えてください。
)

func main() {
	// gRPCサーバーの起動
	grpcServer := grpc.NewServer()

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
		Addr:    ":8080",
		Handler: mux,
	}
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalf("failed to serve HTTP: %v", err)
	}
}
