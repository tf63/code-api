#!/bin/bash
script='query {
  languageCodes(input: {
    languageId: "2",
    startRow: 0,
    endRow: 100,
    offset: 0,
    limit: 2
  }){
    languageCodeId,
    content,
    nrow,
    createdAt,
    languageId
  }
  frameworkCodes(input: {
    toolId: "215",
    startRow: 0,
    endRow: 100,
    offset: 0,
    limit: 2
  }){
    frameworkCodeId,
    content,
    nrow,
    createdAt,
    toolId
  }
  
  patternCodes(input: {
    patternId: "12",
    languageId: "12",
    startRow: 0,
    endRow: 100,
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
    algorithmId: "13",
    languageId: "8",
    startRow: 0,
    endRow: 100,
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
