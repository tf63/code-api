#!/bin/bash
# example
# ----------------------------------------------------------------
# bash bin/loop.sh -n 10 bin/chat/algorithm_code.sh
# ----------------------------------------------------------------

COUNT=10 # initial
while getopts ":n:" opt; do
  case $opt in
    n)
      COUNT="$OPTARG"
      ;;
  esac
done

for ((i=1; i<=$COUNT; i++)) ; do
    source $3
    sleep 30 # 30秒待機
done