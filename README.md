# Tab-commit-gpt

The simplest way to generate commit messages — just press `<Tab>` after `git commit -m`.

_"No new commands. No aliases. No extra steps. Just press Tab."_

While many commit message generators offer a wide range of commands and interfaces, **Tab-commit-gpt focuses on extreme simplicity and usability.**

## Features

- **One-key commit message generation**: Press `<Tab>` during `git commit -m` to instantly generate a commit message.
- **GPT-powered suggestions**: Automatically analyzes your staged Git changes and generates concise, conventional commit messages.
- **Context-aware autocompletion**: If you've already partially typed a commit message (e.g., `git commit -m "fix: ini`), Tab-commit-gpt continues from where you left off, intelligently.
- **Customizable prompts and templates** via environment variables.
- Supports multi-language commit message generation.
- Supports configurable commit message templates.

## Usage

### 1. Generate a message from scratch

```plaintext
# git add "<your files>"
git commit -m
#             ↑ at this point, press Tab
```

Tab-commit-gpt will:

- Read your staged changes.
- Call the GPT API to analyze them.
- Autocomplete the message in your terminal.

### 2. Continue an existing message

```plaintext
# example to generate a commit message beginning with "feat: add"
git commit -m "feat: add
#                        ↑ at this point, press Tab
```

### Prerequisites

- Shell: **Zsh**
- **zsh-autosuggestions** plugin must be installed and configured.
- OpenAI account and API key (charges may apply).
  - You can get your API key from [OpenAI Platform](https://platform.openai.com/settings/).

### 1. Installation

```
brew tap devinjeon/tab-commit-gpt https://github.com/devinjeon/tab-commit-gpt
brew install tab-commit-gpt
```

### 2. Enable the Tab-commit-gpt auto-completion

```bash
echo "source \"$(brew --prefix tab-commit-gpt)/scripts/tab-commit-gpt.sh\"" >> ~/.zshrc
```

## Configuration

You can customize Tab-commit-gpt's behavior using the following environment variables(see `tab-commit-env.sample`):

| Variable                         | Description                                                  | Default                                                       |
| -------------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------- |
| `OPENAI_API_KEY`                 | Your OpenAI API key                                          | **(required)**                                                |
| `OPENAI_MODEL`                   | OpenAI model to use (e.g., `gpt-4o`, `gpt-3.5-turbo`)        | `gpt-4o`                                                      |
| `OPENAI_API_URL`                 | URL for the OpenAI API endpoint                              | `https://api.openai.com/v1/chat/completions`                  |
| `OPENAI_TEMPERATURE`             | Sampling temperature (0.0–1.0, higher = more creative)       | `0.7`                                                         |
| `OPENAI_MAX_TOKENS`              | Maximum number of tokens in response                         | `1024`                                                        |
| `TAB_COMMIT_GPT_LANGUAGE`        | Language for the generated commit message (e.g., `en`, `ko`) | `en`                                                          |
| `TAB_COMMIT_GPT_SOFT_MAX_LENGTH` | Preferred character limit for commit message                 | `60`                                                          |
| `TAB_COMMIT_GPT_HARD_MAX_LENGTH` | Absolute max length for the commit message                   | `80`                                                          |
| `TAB_COMMIT_GPT_TEMPLATE`        | Commit style guide / format (e.g., Conventional Commits)     | `feat\|fix\|chore\|refactor\|test\|style\|docs\|...: message` |
| `TAB_COMMIT_GPT_INSTRUCTIONS`    | Detailed GPT instructions for formatting and tone            | One-liner, no period, clear and concise                       |
| `TAB_COMMIT_GPT_PROMPT_TEMPLATE` | Full prompt template used for GPT request                    | Built-in default prompt template                              |
