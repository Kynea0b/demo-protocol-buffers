#!/bin/bash

# 環境変数の設定 (例: .env ファイルを読み込む場合)
set -a
if [ -f .env ]; then
  source .env
fi
set +a

## grpcurl を使用したリクエスト
#grpcurl -plaintext -v -H "authorization: Bearer $TOKEN" -d "{\"user_id\": \"$USER_ID\"}" localhost:50051 account.AccountService/GetUserInfo
#

## -- grpc --
### Register
grpcurl -plaintext -d '{"username":"foo","password":"password123","email":"foo@example.com"}' localhost:50051 account.AccountService.RegisterUser

### Login
grpcurl -plaintext -d '{"username":"foo","password":"password123"}' localhost:50051 account.AccountService.LoginUser

### Get User Info
grpcurl -plaintext -v -H "authorization: Bearer $TOKEN" -d "{\"user_id\": \"$USER_ID\"}" localhost:50051 account.AccountService/GetUserInfo
