#!/bin/bash

# ユーザー名を引数で指定
if [ $# -ne 1 ]; then
  echo "Usage: $0 <username>"
  exit 1
fi
USERNAME=$1

# パスワードとメールアドレスをユーザー名に基づいて設定
PASSWORD="password123"
EMAIL="${USERNAME}@example.com"

# 環境変数の設定 (例: .env ファイルを読み込む場合)
set -a
if [ -f .env ]; then
  source .env
fi
set +a

## grpcurl を使用したリクエスト
#grpcurl -plaintext -v -H "authorization: Bearer $TOKEN" -d "{\"user_id\": \"$USER_ID\"}" localhost:50051 account.AccountService/GetUserInfo

## -- grpc --
### Register
grpcurl -plaintext -d "{\"username\":\"$USERNAME\",\"password\":\"$PASSWORD\",\"email\":\"$EMAIL\"}" localhost:50051 account.AccountService.RegisterUser

### 異常系テスト
#grpcurl -plaintext -d '{"username":"foo","password":"password123","email":"foo@example.com"}' localhost:50051 account.AccountService.RegisterUser

# Login
result=$(grpcurl -plaintext -d "{\"username\":\"$USERNAME\",\"password\":\"$PASSWORD\"}" localhost:50051 account.AccountService.LoginUser | jq -r '.')

# userId と token の値を取得
user_id=$(echo "$result" | jq -r '.userId')
token=$(echo "$result" | jq -r '.token')

# .env ファイルを更新
if [ -f ../.env ]; then
  sed -i "" "s/^USER_ID=.*/USER_ID=\"$user_id\"/" ../.env
  sed -i "" "s/^TOKEN=.*/TOKEN=\"$token\"/" ../.env
  source ../.env # 環境変数を再読み込み
else
  echo "USER_ID=\"$user_id\"" >> ../.env
  echo "TOKEN=\"$token\"" >> ../.env
  source ../.env # 環境変数を再読み込み
fi

USER_ID=${user_id}
TOKEN=${token}

echo "USER_ID and TOKEN updated in ../.env file."
echo "USER_ID: $USER_ID"
echo "TOKEN: $TOKEN"

# Get User Info
user_info=$(grpcurl -plaintext -H "authorization: Bearer $TOKEN" -d "{\"user_id\": \"$USER_ID\"}" localhost:50051 account.AccountService/GetUserInfo | jq -r '.')

# テスト結果の判定
expected_output="{\"username\":\"$USERNAME\",\"email\":\"$EMAIL\"}"

if [ "$user_info" == "$expected_output" ]; then
  echo "Test Passed!"
  echo "$user_info"
else
  echo "Test Failed!"
  echo "Expected: $expected_output"
  echo "Actual: $user_info"
  exit 1
fi