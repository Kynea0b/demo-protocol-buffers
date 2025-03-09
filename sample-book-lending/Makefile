# root/pkg/grpc直下に作るための命令
gen_code:
	protoc -I ./proto \
    --go_out=./pkg/grpc --go_opt paths=source_relative \
    --go-grpc_out ./pkg/grpc --go-grpc_opt paths=source_relative \
    --grpc-gateway_out ./pkg/grpc --grpc-gateway_opt paths=source_relative \
	./proto/*.proto

gen_doc:
	protoc -I ./proto --doc_out=html,index.html:./doc proto/*.proto

gen_swagger:
	protoc  -I ./proto --openapiv2_out=./doc --openapiv2_opt=disable_service_tags=true,output_format=yaml ./proto/*.proto

start_server:
	cd cmd/server && go run main.go

# Borrow Book
borrow_book_with_grpc:
	grpcurl -plaintext -d '{"book_id":"123","user_id":"xeeddfahwiophg"}' localhost:50051 library.LibraryService.BorrowBook

borrow_book_with_restapi:
	curl -X POST -H "Content-Type: application/json" -d '{"book_id":"123","user_id":"xeeddfahwiophg"}' http://localhost:8080/v1/info/book


fmt:
	go fmt ./...

