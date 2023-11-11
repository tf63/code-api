#!/bin/bash
# ---------------------------------------------------------------------
# OpenAI Playground https://platform.openai.com/playground
# OPENAI_KEYの取得 https://platform.openai.com/account/api-keys
# ---------------------------------------------------------------------
# dependency
#   jq (`brew install jq``)
#   gshuf (`brew install coreutils`)
# ---------------------------------------------------------------------


# ---------------------------------------------------------------------
mkdir -p bin/out bin/data # input/outputの退避先のディレクトリを作成
source bin/init.sh

function get_random() {
    local array=("$@")
    local array_length=${#array[@]}
    local random_index=$((RANDOM % array_length))

    echo ${array[$random_index]}
}

TOOLS=("${!TOOL_ID_MAP[@]}")
TOOL=$(get_random "${TOOLS[@]}")
TOOL_ID=${TOOL_ID_MAP[${TOOL}]}
echo "tool: ${TOOL}, tool_id: ${TOOL_ID}"

LANGUAGE=`echo ${TOOL} | cut -d'-' -f1`
echo "language: ${LANGUAGE}"
FRAMEWORK=`echo ${TOOL} | cut -d'-' -f2`
echo "framework: ${FRAMEWORK}"

OUT_PATH=bin/data/framework_code.csv
if ! [ -e $OUT_PATH ]; then
    echo "nrow,tool_id,content" > $OUT_PATH
fi

cat bin/prompt/framework_code.json | \
    jq ".messages[-1] |= .+ {\"role\": \"user\", \"content\": \"language: ${LANGUAGE}, framework: ${FRAMEWORK}\"}" \
        > bin/input/framework_code.json # 入力を作成

# リクエストを送信
curl https://api.openai.com/v1/chat/completions \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer ${OPENAI_KEY}" \
    -X POST -d @bin/input/framework_code.json > bin/out/framework_code.json

# レスポンスからデータを取得
RESPONSE_DATA=`jq '.choices[0].message.content' bin/out/framework_code.json`
NROW=$(grep -o -n '\\n' <<< "${RESPONSE_DATA}" | wc -l)
NROW=$((NROW+1))
CONTENT="${RESPONSE_DATA:1:${#RESPONSE_DATA}-2}"
CONTENT=${CONTENT#*<code>}
CONTENT=${CONTENT%</code>*}
echo "save to ${OUT_PATH}"
# ---------------------------------------------------------------------
