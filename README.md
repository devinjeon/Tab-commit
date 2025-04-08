# Tabmit

The simplest way to generate commit messages — just press `<Tab>` after `git commit -m`.

*"No new commands. No aliases. No extra steps. Just press Tab."*

While many commit message generators offer a wide range of commands and interfaces, **Tabmit focuses on extreme simplicity and usability.**

## Features

- **One-key commit message generation**: Press `Tab` during `git commit -m` to instantly generate a commit message.
- **GPT-powered suggestions**: Automatically analyzes your staged Git changes and generates concise, conventional commit messages.
- **Context-aware autocompletion**: If you've already partially typed a commit message (e.g., `git commit -m "fix: ini`), Tabmit continues from where you left off, intelligently.
- **Customizable prompts and templates** via environment variables.
- Supports multi-language commit message generation.
- Supports configurable commit message templates.

##  Usage

### 1. Generate a message from scratch

```plaintext
# git add "<your files>"
git commit -m
#             ↑ at this point, press Tab
```

Tabmit will:
- Read your staged changes.
- Call the GPT API to analyze them.
- Autocomplete the message in your terminal.

### 2. Continue an existing message

```plaintext
# example to generate a commit message beginning with "feat: add"
git commit -m "feat: add
#                        ↑ at this point, press Tab
```

## Installation

(TODO: add package manager support)

### 1. Prerequisites

(TODO: add various shell support)

- Shell: **Zsh**
- **zsh-autosuggestions** plugin must be installed and configured.
- OpenAI account and API key (charges may apply).
   - You can get your API key from [OpenAI Platform](https://platform.openai.com/settings/).

### 2. Set environment variables

```bash
cp tabmit-env.example tabmit-env # set your configuration
echo "source $PWD/tabmit-env" >> ~/.zshrc
```

### 3. Build the binary

Clone the repository and run:

```bash
go build -o tabmit .
echo "export PATH=\"$PWD:\$PATH\"" >> ~/.zshrc
```

### 2. Enable the Tabmit auto-completion

```bash
echo "source $PWD/tabmit-autocomplete.zsh" >> ~/.zshrc
```

## Configuration

You can customize Tabmit's behavior using the following environment variables(see `tabmit-env.sample`):

| Variable                 | Description                                                  | Default                                                       |
| ------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------- |
| `OPENAI_API_KEY`         | Your OpenAI API key                                          | **(required)**                                                |
| `OPENAI_MODEL`           | OpenAI model to use (e.g., `gpt-4o`, `gpt-3.5-turbo`)        | `gpt-4o`                                                      |
| `OPENAI_API_URL`         | URL for the OpenAI API endpoint                              | `https://api.openai.com/v1/chat/completions`                  |
| `OPENAI_TEMPERATURE`     | Sampling temperature (0.0–1.0, higher = more creative)       | `0.7`                                                         |
| `OPENAI_MAX_TOKENS`      | Maximum number of tokens in response                         | `1024`                                                        |
| `TABMIT_LANGUAGE`        | Language for the generated commit message (e.g., `en`, `ko`) | `en`                                                          |
| `TABMIT_SOFT_MAX_LENGTH` | Preferred character limit for commit message                 | `60`                                                          |
| `TABMIT_HARD_MAX_LENGTH` | Absolute max length for the commit message                   | `80`                                                          |
| `TABMIT_TEMPLATE`        | Commit style guide / format (e.g., Conventional Commits)     | `feat\|fix\|chore\|refactor\|test\|style\|docs\|...: message` |
| `TABMIT_INSTRUCTIONS`    | Detailed GPT instructions for formatting and tone            | One-liner, no period, clear and concise                       |
| `TABMIT_PROMPT_TEMPLATE` | Full prompt template used for GPT request                    | Built-in default prompt template                              |
