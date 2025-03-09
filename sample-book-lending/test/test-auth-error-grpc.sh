#!/bin/bash

# テスト対象のエンドポイント (gRPC サーバーのアドレス)
GRPC_SERVER="localhost:50051"

# テスト対象のメソッド
METHOD="account.AccountService/GetUserInfo"

# テスト用のユーザーID
USER_ID="test_user_id"

# 1. メタデータがない場合
echo "Testing: Metadata not provided"
grpcurl -plaintext -d "{\"user_id\": \"${USER_ID}\"}" ${GRPC_SERVER} ${METHOD} 2> error.log
if [ $? -ne 0 ]; then
  if grep -q "Unauthenticated" error.log; then
    echo "Test Passed: Metadata not provided"
  else
    echo "Test Failed: Expected Unauthenticated, but got error: $(cat error.log)"
    exit 1
  fi
else
  echo "Test Failed: Expected Unauthenticated, but got success"
  exit 1
fi

# 2. Authorization ヘッダーがない場合
echo "Testing: Authorization header not provided"
grpcurl -plaintext -d "{\"user_id\": \"${USER_ID}\"}" ${GRPC_SERVER} ${METHOD} 2> error.log
if [ $? -ne 0 ]; then
  if grep -q "Unauthenticated" error.log; then
    echo "Test Passed: Authorization header not provided"
  else
    echo "Test Failed: Expected Unauthenticated, but got error: $(cat error.log)"
    exit 1
  fi
else
  echo "Test Failed: Expected Unauthenticated, but got success"
  exit 1
fi

# 3. 不正なトークンの場合
echo "Testing: Invalid token"
grpcurl -plaintext -H "authorization: Bearer invalid_token" -d "{\"user_id\": \"${USER_ID}\"}" ${GRPC_SERVER} ${METHOD} 2> error.log
if [ $? -ne 0 ]; then
  if grep -q "Unauthenticated" error.log; then
    echo "Test Passed: Invalid token"
  else
    echo "Test Failed: Expected Unauthenticated, but got error: $(cat error.log)"
    exit 1
  fi
else
  echo "Test Failed: Expected Unauthenticated, but got success"
  exit 1
fi

## 4. 無効なトークンの場合 (例: 有効期限切れ)
#echo "Testing: Expired token"
## 有効期限切れのトークンを生成 (jwt コマンドなどを使用)
#EXPIRED_TOKEN="f9d5278d-bf50-4efb-b7b4-c54f2e8da874" # 置き換えてください
#result=$(grpcurl -plaintext -H "authorization: Bearer ${EXPIRED_TOKEN}" -d "{\"user_id\": \"${USER_ID}\"}" ${GRPC_SERVER} ${METHOD} | jq -r '.')
#if [[ "$result" == *"Unauthenticated"* ]]; then
#  echo "Test Passed: Expired token"
#else
#  echo "Test Failed: Expected Unauthenticated, but got: $result"
#  exit 1
#fi
#

#
## 6. 不正なユーザーID の場合 (例: ユーザーID が数値でない)
#echo "Testing: Invalid user ID in token"
## ユーザーID が数値でないトークンを生成 (jwt コマンドなどを使用)
#INVALID_USER_ID_TOKEN="<invalid_user_id_token>" # 置き換えてください
#result=$(grpcurl -plaintext -H "authorization: Bearer ${INVALID_USER_ID_TOKEN}" -d "{\"user_id\": \"${USER_ID}\"}" ${GRPC_SERVER} ${METHOD} | jq -r '.')
#if [[ "$result" == *"Unauthenticated"* ]]; then
#  echo "Test Passed: Invalid user ID in token"
#else
#  echo "Test Failed: Expected Unauthenticated, but got: $result"
#  exit 1
#fi

echo "All tests passed!"