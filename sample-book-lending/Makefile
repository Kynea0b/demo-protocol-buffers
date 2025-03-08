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


# Register User
register_user_with_grpc:
	grpcurl -plaintext -d '{"username":"testuserhoge","password":"password123","email":"testhoge@example.com"}' localhost:50051 account.AccountService.RegisterUser

#todo not found
register_user_with_api:
	curl -X POST -H "Content-Type: application/json" -d '{"username":"testuser","password":"password123","email":"test@example.com"}' http://localhost:8080/v1/accounts/register

# Login
login:
	grpcurl -plaintext -d '{"username":"testuser","password":"password123"}' localhost:50051 account.AccountService.LoginUser

login_user_with_api:
	curl -X POST -H "Content-Type: application/json" -d '{"username":"testuser","password":"password123"}' http://localhost:8080/v1/accounts/login

get_user_info:
	grpcurl -plaintext -d '{"user_id":"a6480bd2-ef6f-430f-ba3c-649fec8cf21e"}' localhost:50051 account.AccountService.GetUserInfo

get_user_info_with_api:
	curl http://localhost:8080/v1/accounts/"427c86e3-086e-4fc1-8a88-9a8d11e8a6e5"


fmt:
	go fmt ./...

