#!/bin/sh

PWD=$(pwd)
SESSION=$(basename "$PWD")

tmux new-session -s "$SESSION" -d
tmux send-keys -t "$SESSION:0.0" '(cd gqlgen && gow -e=gql,yml run github.com/99designs/gqlgen generate)' C-m

tmux split-window -v -t "$SESSION"
tmux select-layout even-vertical   # to avoid 'no space for new pane' 
tmux send-keys -t "$SESSION:0.1" '(cd gqlgen && gow -e=go,json run . server)' C-m

tmux split-window -v -t "$SESSION"
tmux select-layout even-vertical   # to avoid 'no space for new pane' 
tmux send-keys -t "$SESSION:0.2" '(cd next && npm run dev)' C-m

tmux split-window -v -t "$SESSION" 
tmux select-layout even-vertical   # to avoid 'no space for new pane' 
tmux send-keys -t "$SESSION:0.3" '(cd next && npm run codegen)' C-m

# vs code's bottom panel shows TypeScript errors as "Problems"
# tmux split-window -v -t "$SESSION" 
# tmux select-layout even-vertical   # to avoid 'no space for new pane' 
# tmux send-keys -t "$SESSION:0.4" '(cd next && npm run compile)' C-m

# vs code's bottom panel shows TypeScript errors as "Problems"
# tmux split-window -v -t "$SESSION" 
# tmux select-layout even-vertical   # to avoid 'no space for new pane' 
# tmux send-keys -t "$SESSION:0.4" '(cd next && npm run lint-watch)' C-m

tmux split-window -v -t "$SESSION" 
tmux select-layout even-vertical   # to avoid 'no space for new pane' 
tmux send-keys -t "$SESSION:0.4" '(cd gqlgen && gow -e=go,json,gql -c -v test ./...)' C-m

# open editors
tmux split-window -v -t "$SESSION" 
tmux select-layout even-vertical   # to avoid 'no space for new pane' 
tmux send-keys -t "$SESSION:0.5" 'code gqlgen && code next' C-m

tmux attach -t "$SESSION"
