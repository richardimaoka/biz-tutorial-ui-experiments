#!/bin/sh

set -e

BASEDIR="./"
for j in 0 1 # 2 3 4 5 6 7 8 9
do
  for i in 0 1 2 3 4 5 6 7 8 9
  do
    FILENAME="$BASEDIR/action$j$i.json"
    if [ ! -f "$FILENAME" ]
    then
      echo "" > "$FILENAME"
    fi
  done
done

#TODO: create a script to generate state.json from , maybe a Go script
# directory structure
#   data
#     |- step01
#          |- action.json
#          |- state.json 
#     |- step02
#          |- action.json
#          |- state.json //auto generated from prev state + action