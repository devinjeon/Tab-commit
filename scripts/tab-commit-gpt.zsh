#!/usr/bin/env zsh

_tab-commit-gpt() {
  if [[ $BUFFER == git\ commit\ -m\ * ]]; then
    # Check OPENAI_API_KEY
    if [[ -z "$OPENAI_API_KEY" ]]; then
      zle -M "The OPENAI_API_KEY is not set. Please refer to the guide for setup instructions: https://github.com/devinjeon/Tab-commit-gpt#configuration"
      return 0
    fi

    # Generate completion
    completion=$(tab-commit-gpt "$BUFFER" 2>/dev/null)

    if [[ -n "$completion" ]]; then
      local new_msg="${completion}"

      # Replace the whole buffer and reposition the cursor
      BUFFER="$completion"
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
zle -N _tab-commit-gpt
bindkey '^I' _tab-commit-gpt
