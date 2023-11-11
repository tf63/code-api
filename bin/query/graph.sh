#!/bin/bash
script='query {
  programCodes(input: {
    toolId: "17",
    startRow: 0,
    endRow: 100,
    offset: 0,
    limit: 2
  }){
    programCodeId,
    content,
    nrow,
    createdAt,
    toolId
  }
  
  patternCodes(input: {
    patternId: "1",
    languageId: "1",
    startRow: 0,
    endRow: 10,
    offset: 0,
    limit: 2
  }){
    patternCodeId,
    content,
    nrow,
    createdAt,
    patternId,
    languageId
  }
  
  algorithmCodes(input: {
    algorithmId: "1",
    languageId: "1",
    startRow: 0,
    endRow: 10,
    offset: 0,
    limit: 2
  }){
    algorithmCodeId,
    content,
    nrow,
    createdAt,
    algorithmId,
    languageId
  }
  
  languages{
    languageId,
    name
  }
  
  frameworks{
    frameworkId,
    name
  }
  
  patterns {
    patternId,
    name
  }
  
  algorithms {
    algorithmId,
    name
  }
}'

# シングルクォートで囲まれた変数内のダブルクォートをエスケープ
script="${script//\"/\\\"}"
script="$(echo $script)"

# CURLコマンドでJSONリクエストを送信
curl "http://localhost:8080/query" \
    -H "Content-Type: application/json" \
    -X POST -d "{\"query\": \"$script\"}"
