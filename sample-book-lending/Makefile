# root/pkg/grpc直下に作るための命令
gen_code:
	protoc -I ./proto \
    --go_out=./pkg/grpc --go_opt paths=source_relative \
    --go-grpc_out ./pkg/grpc --go-grpc_opt paths=source_relative \
    --grpc-gateway_out ./pkg/grpc --grpc-gateway_opt paths=source_relative \
	./proto/*.proto

gen-doc:
	protoc -I ./proto --doc_out=html,index.html:./doc proto/*.proto

start_server:
	cd cmd/server && go run main.go

call_sendborrow_method_of_the_service:
	grpcurl -plaintext -d '{"account":{ "name": "Tanaka"},"book":{ "title": "赤毛のあん"}}' localhost:8080 myapp.LendingBooksService.SendBorrow

request:
	curl -X POST -k http://localhost:8090/v1/accountinfo/books -d '{"name": " Proxy Anpanman"}'

