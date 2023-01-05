#!/bin/sh

set -e

BASEDIR="data/tutorial1"
for j in 0 1 2 3 4 5 6 7 8 9
do
  for i in 0 1 2 3 4 5 6 7 8 9
  do
    FILENAME="$BASEDIR/step$j$i/state.json"
    if [ -f $FILENAME ]
    then
      mv "$FILENAME" "$BASEDIR/step$j$i.json"
    fi
  done
done