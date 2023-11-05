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
script="$(echo $script)" #改行を消して1行にする、改行あるとJSON parsingエラーが返ってくる

# CURLコマンドでJSONリクエストを送信
curl -X POST "http://localhost:8080/query" \
    -H "Content-Type: application/json" \
    -d "{\"query\": \"$script\"}"