#!/bin/bash

# 環境変数の設定 (例: .env ファイルを読み込む場合)
set -a
if [ -f .env ]; then
  source .env
fi
set +a

## curl を使用したリクエスト
#curl -v -H "Authorization: Bearer $TOKEN" http://localhost:35635/v1/accounts/$USER_ID
### Register
curl -X POST -H "Content-Type: application/json" -d '{"username":"bar","password":"password123","email":"bar@example.com"}' http://localhost:35635/v1/accounts/register

### Login
curl -X POST -H "Content-Type: application/json" -d '{"username":"bar","password":"password123"}' http://localhost:35635/v1/accounts/login

### Get User Info
curl -H "Authorization: Bearer $TOKEN" http://localhost:35635/v1/accounts/$USER_ID
