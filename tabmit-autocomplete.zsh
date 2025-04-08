#!/usr/bin/env zsh

_zsh_autosuggest_strategy__tabmit() {
  local BUFFER_STR=${BUFFER}

  if [[ "$BUFFER_STR" =~ ^git\ commit\ -m\ \"([^\"]*)$ ]]; then
    local partial_msg="${match[1]}"
    local suggestion

    suggestion=$(tabmit "$partial_msg" 2>/dev/null)

    if [[ -n "$suggestion" ]]; then
      suggestion="${suggestion#${partial_msg}}"
      echo "$suggestion\""
    fi
  fi
}
export ZSH_AUTOSUGGEST_STRATEGY=(_tabmit "$ZSH_AUTOSUGGEST_STRATEGY[@]")

_tabmit() {
  if [[ $BUFFER == git\ commit\ -m\ * ]]; then
    local raw_input="${BUFFER#git commit -m }"
    local msg_prefix
    local quoted=false

    # Remove starting quote if present
    if [[ $raw_input == \"* ]]; then
      quoted=true
      msg_prefix="${raw_input#\"}"
    else
      msg_prefix="$raw_input"
    fi

    echo "msg_prefix: $msg_prefix" >> /tmp/gpt_commit_msg.log
    # Generate completion
    local completion
    completion=$(tabmit "$msg_prefix" 2>/dev/null)
    echo "completion: $completion" >> /tmp/gpt_commit_msg.log

    if [[ -n "$completion" ]]; then
      # local new_msg="${msg_prefix}${completion}"
      local new_msg="${completion}"
      local new_cmd="git commit -m \"${new_msg}\""
      echo "new_cmd: $new_cmd" >> /tmp/gpt_commit_msg.log

      # Replace the whole buffer and reposition the cursor
      BUFFER="$new_cmd"
      CURSOR=${#BUFFER}
      zle reset-prompt
    else
      zle -M "Error generating commit message."
    fi

    return 0
  else
    zle expand-or-complete
    # _git
  fi
}

zle -N _tabmit
bindkey '^I' _tabmit
