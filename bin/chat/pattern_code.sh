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

LANGUAGES=("${!LANGUAGE_ID_MAP[@]}")
LANGUAGE=$(get_random "${LANGUAGES[@]}")
LANGUAGE_ID=${LANGUAGE_ID_MAP[${LANGUAGE}]}
echo "language: ${LANGUAGE}, language_id: ${LANGUAGE_ID}"

PATTERNS=("${!PATTERN_ID_MAP[@]}")
PATTERN=$(get_random "${PATTERNS[@]}")
PATTERN_ID=${PATTERN_ID_MAP[${PATTERN}]}
echo "pattern: ${PATTERN}, pattern_id: ${PATTERN_ID}"

OUT_PATH=bin/data/pattern_code.csv
if ! [ -e $OUT_PATH ]; then
    echo "nrow,language_id,pattern_id,content" > $OUT_PATH
fi

cat bin/prompt/pattern_code.json | \
    jq ".messages[-1] |= .+ {\"role\": \"user\", \"content\": \"language: ${LANGUAGE}, pattern: ${PATTERN}\"}" \
        > bin/input/pattern_code.json # 入力を作成

# リクエストを送信
curl https://api.openai.com/v1/chat/completions \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer ${OPENAI_KEY}" \
    -X POST -d @bin/input/pattern_code.json > bin/out/pattern_code.json

# レスポンスからデータを取得
RESPONSE_DATA=`jq '.choices[0].message.content' bin/out/pattern_code.json`
NROW=$(grep -o -n '\\n' <<< "${RESPONSE_DATA}" | wc -l)
NROW=$((NROW+1))
CONTENT="${RESPONSE_DATA:1:${#RESPONSE_DATA}-2}"
CONTENT=${CONTENT#*<code>}
CONTENT=${CONTENT%</code>*}
echo "${NROW},${LANGUAGE_ID},${PATTERN_ID},\"${CONTENT}\"" >> $OUT_PATH
echo "save to ${OUT_PATH}"
# ---------------------------------------------------------------------
