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

register_book:
	curl -X POST -k http://localhost:8090/v1/info/book -d '{"num": 3, "title": "七つの海のティコ"}'


# todo use body
register_account:
	curl -X POST -k http://localhost:8090/v1/register/"swimming" -d '{"name": "nyannnn", "mail": "hogehogehoge@gmail.com"}'


sendborrow1:
	grpcurl -plaintext -d '{"account":{ "name": "Tanaka"},"book":{ "title": "赤毛のアン"}}' localhost:8080 myapp.LendingBooksService.SendBorrow

sendborrow2:
	grpcurl -plaintext -d '{"account":{ "name": "Tanaka"},"book":{ "title": "七つの海のティコ"}}' localhost:8080 myapp.LendingBooksService.SendBorrow

get_lendinginfo:
	curl -X GET http://localhost:8090/v1/info/book/"七つの海のティコ"

get_borrowedtime:
	curl -X GET http://localhost:8090/v1/info/account/"Tanaka"

hello:
	curl -X GET http://localhost:8090/hello/"foobarbaz"
