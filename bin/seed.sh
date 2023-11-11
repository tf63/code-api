#!/bin/bash

mkdir -p bin/data
source bin/seed/init.sh
# ----------------------------------------------------------------
# language
# ----------------------------------------------------------------
echo "language_id,name" > bin/data/language.csv
for language in "${!LANGUAGE_ID_MAP[@]}"; do
    echo "${LANGUAGE_ID_MAP[$language]},${language}" >> bin/data/language.csv
done
# ----------------------------------------------------------------
# framework
# ----------------------------------------------------------------
echo "framework_id,name" > bin/data/framework.csv
for framework in "${!FRAMEWORK_ID_MAP[@]}"; do
    echo "${FRAMEWORK_ID_MAP[$framework]},${framework}" >> bin/data/framework.csv
done
# ----------------------------------------------------------------
# tool
# ----------------------------------------------------------------
echo "tool_id,language_id,framework_id" > bin/data/tool.csv
for tool in "${!TOOL_ID_MAP[@]}"; do
    language=`echo ${tool} | cut -d'-' -f1`
    framework=`echo ${tool} | cut -d'-' -f2`
    echo "${TOOL_ID_MAP[$tool]},${LANGUAGE_ID_MAP[${language}]},${FRAMEWORK_ID_MAP[${framework}]}" >> bin/data/tool.csv
done
# ----------------------------------------------------------------
# algorithm
# ----------------------------------------------------------------
echo "algorithm_id,name" > bin/data/algorithm.csv
for algorithm in "${!ALGORITHM_ID_MAP[@]}"; do
    echo "${ALGORITHM_ID_MAP[$algorithm]},${algorithm}" >> bin/data/algorithm.csv
done
# ----------------------------------------------------------------
# pattern
# ----------------------------------------------------------------
echo "pattern_id,name" > bin/data/pattern.csv
for pattern in "${!PATTERN_ID_MAP[@]}"; do
    echo "${PATTERN_ID_MAP[$pattern]},${pattern}" >> bin/data/pattern.csv
done