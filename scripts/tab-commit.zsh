#!/usr/bin/env zsh

_tab-commit() {
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

    # Generate completion
    local completion
    completion=$(tab-commit "$msg_prefix" 2>/dev/null)

    if [[ -n "$completion" ]]; then
      local new_msg="${completion}"
      local new_cmd="git commit -m \"${new_msg}\""

      # Replace the whole buffer and reposition the cursor
      BUFFER="$new_cmd"
      CURSOR=${#BUFFER}
      zle reset-prompt
      # remove autosuggest if present
      zle autosuggest-clear
    else
      zle -M "Error generating commit message."
    fi

    return 0
  else
    zle expand-or-complete
  fi
}

# This is a temporary workaround for auto completion behavior instead of using compadd
# compadd is not working as expected when input commit message is not closed with a double quote
zle -N _tab-commit
bindkey '^I' _tab-commit
