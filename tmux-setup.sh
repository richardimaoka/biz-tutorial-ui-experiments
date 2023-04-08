#!/bin/sh

PWD=$(pwd)
SESSION=$(basename "$PWD")

tmux new-session -s "$SESSION" -d

tmux send-keys -t "$SESSION:0.0" '(cd gqlgen && gow -e=gql run github.com/99designs/gqlgen generate)' C-m
tmux split-window -v -t "$SESSION"
tmux select-layout even-vertical   # to avoid 'no space for new pane' 

tmux send-keys -t "$SESSION:0.1" '(cd gqlgen && gow -e=go,json run main.go)' C-m
tmux split-window -v -t "$SESSION"
tmux select-layout even-vertical   # to avoid 'no space for new pane' 

tmux send-keys -t "$SESSION:0.2" '(cd next && npm run dev)' C-m
tmux split-window -v -t "$SESSION" 
tmux select-layout even-vertical   # to avoid 'no space for new pane' 

tmux send-keys -t "$SESSION:0.3" '(cd next && npm run codegen)' C-m
tmux split-window -v -t "$SESSION" 
tmux select-layout even-vertical   # to avoid 'no space for new pane' 

tmux send-keys -t "$SESSION:0.4" '(cd next && npm run compile)' C-m
tmux split-window -v -t "$SESSION" 
tmux select-layout even-vertical   # to avoid 'no space for new pane' 

# TODO: enable watch mode for next lint
# tmux send-keys -t "$SESSION:0.5" '(cd next && npm run lint)' C-m
# tmux split-window -v -t "$SESSION" 
# tmux select-layout even-vertical   # to avoid 'no space for new pane' 

# open editors
tmux send-keys -t "$SESSION:0.5" '(cd gqlgen && code .) && (cd next && code .)' C-m

tmux attach -t "$SESSION"
