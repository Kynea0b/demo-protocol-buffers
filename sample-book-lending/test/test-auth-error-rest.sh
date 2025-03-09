#!/bin/bash

# テスト対象のエンドポイント (REST API サーバーのアドレス)
REST_API_SERVER="http://localhost:8080/v1/accounts"

# テスト対象のユーザーID
USER_ID="test_user_id"

# 1. Authorization ヘッダーがない場合
echo "Testing: Authorization header not provided"
response=$(curl -s -o /dev/null -w "%{http_code}" ${REST_API_SERVER}/${USER_ID})
if [ "$response" -eq 401 ]; then
  echo "Test Passed: Authorization header not provided"
else
  echo "Test Failed: Expected 401, but got $response"
  exit 1
fi

# 2. 不正なトークンの場合
echo "Testing: Invalid token"
response=$(curl -s -o /dev/null -w "%{http_code}" -H "Authorization: Bearer invalid_token" ${REST_API_SERVER}/${USER_ID})
if [ "$response" -eq 401 ]; then
  echo "Test Passed: Invalid token"
else
  echo "Test Failed: Expected 401, but got $response"
  exit 1
fi

## 3. 有効期限切れのトークンの場合
#echo "Testing: Expired token"
## 有効期限切れのトークンを生成 (jwt コマンドなどを使用)
#EXPIRED_TOKEN="<expired_token>" # 置き換えてください
#response=$(curl -s -o /dev/null -w "%{http_code}" -H "Authorization: Bearer ${EXPIRED_TOKEN}" ${REST_API_SERVER}/${USER_ID})
#if [ "$response" -eq 401 ]; then
#  echo "Test Passed: Expired token"
#else
#  echo "Test Failed: Expected 401, but got $response"
#  exit 1
#fi

# 4. 不正なトークン形式の場合
echo "Testing: Invalid token format"
response=$(curl -s -o /dev/null -w "%{http_code}" -H "Authorization: invalid_token_format" ${REST_API_SERVER}/${USER_ID})
if [ "$response" -eq 401 ]; then
  echo "Test Passed: Invalid token format"
else
  echo "Test Failed: Expected 401, but got $response"
  exit 1
fi

echo "All tests passed!"