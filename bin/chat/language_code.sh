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

THEMES=("list" "map" "dictonary" "set" "io" "class" \
    "function" "for" "if" "test" "error" "recursive")
THEME=$(get_random ${THEMES[@]})
echo "theme: ${THEME}"

OUT_PATH=bin/data/language_code.csv
if ! [ -e $OUT_PATH ]; then
    echo "nrow,language_id,content" > $OUT_PATH
fi

cat bin/prompt/language_code.json | \
    jq ".messages[-1] |= .+ {\"role\": \"user\", \"content\": \"language: ${LANGUAGE}, theme: ${THEME}\"}" \
        > bin/input/language_code.json # 入力を作成

# リクエストを送信
curl https://api.openai.com/v1/chat/completions \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer ${OPENAI_KEY}" \
    -X POST -d @bin/input/language_code.json > bin/out/language_code.json

# レスポンスからデータを取得
RESPONSE_DATA=`jq '.choices[0].message.content' bin/out/language_code.json`
NROW=$(grep -o -n '\\n' <<< "${RESPONSE_DATA}" | wc -l)
NROW=$((NROW+1))
CONTENT="${RESPONSE_DATA:1:${#RESPONSE_DATA}-2}"
CONTENT=${CONTENT#*<code>}
CONTENT=${CONTENT%</code>*}
echo "${NROW},${LANGUAGE_ID},\"${CONTENT}\"" >> $OUT_PATH
echo "save to ${OUT_PATH}"
# ---------------------------------------------------------------------
