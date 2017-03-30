tmux
`ctrl+b - prefix key`

creating and navigating panes
```
ctrl+b + % - vertcal pane
ctrl+b + " - horizontal pane
ctrl+b + arrow key - to navigate
exit or ctrl+d - to remove pane
```

`ctrl+b + ? - all different shortcuts`

windows - collection of panes
```
ctrl+b + c - to create a new window
ctrl+b n,p - to navigate within windows
```

detaching and attaching to tmux sessions
```
ctrl+b + d - detach tmux
tmux attach - will attach to last detached tmux session
tmux ls - lists tmux sessions
tmux attach -t <session_id>
eg: tmux attach -t 0
```

```
ctrl+b + $ - to name your session
ctrl+b + s - to view sessions within tmux

tmux rename-session -t 2 <new_name> - to rename a session from commandline
tmux new -s <session_name> - to create a session with a name
```


`ctrl+b + z - to toggle zoom. Enters full screen with the current pane`

to resize panes
```
ctrl+b + : - to go into command mode
:resize-pane -D - to move downwards
:resize-pane -D 30 - move 30 cells downward
-U - upward
-R - right
-L - left
```


There is a difference between detach and ending the session in tmux. If
the last tmux pane is closed then tmux session ends there. If we detach
a tmux session (`ctrl+b + d`) we can reattach to it (`tmux attach -t
<session_id>`).

To configure tmux, modify tmux.conf

to modify control key:
```
unbind C-b
set-option -g prefix C-a
bind-key C-a send-prefix
```


to load new config kill the running sessions and do tmux ls or go to
command mode `(ctrl+a + :)`, `source-file ~/.tmux.conf`

to kill tmux session
`tmux kill-session -t <session_name>`


to make scrolling work
`ctrl+a + [ - copy mode and use arrow keys`
to quit copy mode press `q`.

on macOSX copy, paste works on local machine

but on remote shells we need to use tmux's native copy, paste
capabilities. there are Emacs, vim style key bindings.

for vim: go to command mode using ctrl+a + :
`set-window-option mode-keys vi`
then you can navigate using arrow or vim key bindings. press space to
start copying and to quit copy mode press enter. to paste press `ctrl+a +
]`

enable mouse mode in tmux
set mouse on

to put key bindings in ~/.tmux.conf file
```
bind m set -g mouse on
bind M set -g mouse off
```

if we want to have same lay for a project. tmux scripting can be used
for this.

sharing tmux session for pair programming with ssh. Supports builtin.

sharing history from different panes and sessions. Put the following
lines in ~/.bashrc
```
shopt -s histappend
shopt -s histreedit
shopt -s histverify
# ignores both commands that start with a white space and duplicate commands
HISTCONTROL='ignoreboth'
PROMPT_COMMAND="history -a; history -c; history -r; $PROMPT_COMMAND"
```

