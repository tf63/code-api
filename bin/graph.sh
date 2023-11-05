#!/bin/bash
script='query {
  programCodes(input: {
    toolId: "1",
    startRow: 0,
    endRow: 10, 
    offset: 0,
    limit: 2
  }
  ){
    programCodeId,
    content,
    nrow,
    createdAt
  }
}'

# シングルクォートで囲まれた変数内のダブルクォートをエスケープ
script="${script//\"/\\\"}"
script="$(echo $script)"

# CURLコマンドでJSONリクエストを送信
curl "http://localhost:8080/query" \
    -H "Content-Type: application/json" \
    -X POST -d "{\"query\": \"$script\"}"
