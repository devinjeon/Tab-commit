package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/devinjeon/kubectl-gpt/pkg/gpt"
)

const (
	defultMaxTokens       = 1024
	defaultTemperature    = 0.7
	defaultModel          = "gpt-4o"
	defaultAPIURL         = "https://api.openai.com/v1/chat/completions"
	defaultLanguage       = "en"
	defaultSoftMaxLength  = "60"
	defaultHardMaxLength  = "80"
	defaultPromptTemplate = `--- Analyze the following changes ---
{{changes}}

--- Instructions ---
Language: {{language}}
Max commit message length (prefered): {{softMaxLength}}
Max commit message length (required): {{hardMaxLength}}
Instructions:
{{instructions}}
Commit message template:
{{commitMessageTemplate}}`

	defaultInstructions = `* Be clear, simple, and concise
* remove articles if necessary
* Do not include a period at the end of the sentence
* MUST write *only* one-line commit message
* If part of the commit message is already written, continue from where it left off (and prefer follow the commit message template if possible)`
	defaultCommitMessageTemplate = "feat|fix|chore|refactor|test|style|docs|...: some message"
)

func getGitChanges() (string, error) {
	// Check if git is installed
	if _, err := exec.LookPath("git"); err != nil {
		return "", fmt.Errorf("git is not installed: %v", err)
	}

	// Check if the current directory is a git repository
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("not a git repository: %v", err)
	}

	// Get the list of staged files
	cmd = exec.Command("git", "diff", "--cached", "--name-only")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to get git changes: %v", err)
	}
	files := strings.TrimSpace(out.String())
	if files == "" {
		return "", fmt.Errorf("no staged changes found")
	}

	// Optionally include a diff preview
	cmdDiff := exec.Command("git", "diff", "--cached")
	var diffOut bytes.Buffer
	cmdDiff.Stdout = &diffOut
	if err := cmdDiff.Run(); err != nil {
		return "", fmt.Errorf("failed to get git diff: %v", err)
	}

	return fmt.Sprintf("Files:\n%s\n\nDiff:\n%s", files, diffOut.String()), nil
}

func getCompletion(prompt string) (string, error) {
	apiURL := defaultAPIURL
	if apiURLEnv := os.Getenv("OPENAI_API_URL"); apiURLEnv != "" {
		apiURL = apiURLEnv
	}

	model := defaultModel
	if modelEnv := os.Getenv("OPENAI_MODEL"); modelEnv != "" {
		model = modelEnv
	}

	temperature := defaultTemperature
	if temperatureEnv := os.Getenv("OPENAI_TEMPERATURE"); temperatureEnv != "" {
		if temperatureFloat, err := strconv.ParseFloat(temperatureEnv, 64); err == nil {
			temperature = temperatureFloat
		}
	}

	maxTokens := defultMaxTokens
	if maxTokensEnv := os.Getenv("OPENAI_MAX_TOKENS"); maxTokensEnv != "" {
		if maxTokensInt, err := strconv.Atoi(maxTokensEnv); err == nil {
			maxTokens = maxTokensInt
		}
	}

	apiToken := os.Getenv("OPENAI_API_KEY")
	if apiToken == "" {
		return "", fmt.Errorf("OPENAI_API_KEY environment variable is not set")
	}

	systemMessage := "You are a assistant for generating commit messages."
	reqBody := gpt.NewOpenAIRequest(model, temperature, maxTokens, systemMessage, prompt)

	respData, err := gpt.RequestChatGptAPI(apiURL, reqBody, apiToken)
	if err != nil {
		return "", fmt.Errorf("failed to get completion: %v", err)
	}

	if len(respData.Choices) == 0 {
		return "", fmt.Errorf("no completion found in response")
	}

	return strings.TrimSpace(respData.Choices[0].Message.Content), nil
}

func main() {
	softMaxLength := defaultSoftMaxLength
	if softMaxLengthEnv := os.Getenv("TABMIT_SOFT_MAX_LENGTH"); softMaxLengthEnv != "" {
		softMaxLength = softMaxLengthEnv
	}

	hardMaxLength := defaultHardMaxLength
	if hardMaxLengthEnv := os.Getenv("TABMIT_HARD_MAX_LENGTH"); hardMaxLengthEnv != "" {
		hardMaxLength = hardMaxLengthEnv
	}

	language := defaultLanguage
	if languageEnv := os.Getenv("TABMIT_LANGUAGE"); languageEnv != "" {
		language = languageEnv
	}

	instructions := defaultInstructions
	if instructionsEnv := os.Getenv("TABMIT_INSTRUCTIONS"); instructionsEnv != "" {
		instructions = instructionsEnv
	}

	commitMessageTemplate := defaultCommitMessageTemplate
	if commitMessageTemplateEnv := os.Getenv("TABMIT_TEMPLATE"); commitMessageTemplateEnv != "" {
		commitMessageTemplate = commitMessageTemplateEnv
	}

	promptTemplate := defaultPromptTemplate
	if promptTemplateEnv := os.Getenv("TABMIT_PROMPT_TEMPLATE"); promptTemplateEnv != "" {
		promptTemplate = promptTemplateEnv
	}

	changes, err := getGitChanges()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// get all arguments from command line and merge strings to one string
	existingCommitMsg := ""
	prompt := promptTemplate
	if len(os.Args) > 1 {
		existingCommitMsg = strings.Join(os.Args[1:], " ")
		if existingCommitMsg != "" {
			prompt = prompt + `The beginning of the commit message has already been written:
{{existingCommitMsg}}
Please continue the commit message **starting from where it left off**, MUST not modify the existing message`
		}
	}

	prompt = strings.Replace(prompt, "{{changes}}", changes, -1)
	prompt = strings.Replace(prompt, "{{language}}", language, -1)
	prompt = strings.Replace(prompt, "{{softMaxLength}}", softMaxLength, -1)
	prompt = strings.Replace(prompt, "{{hardMaxLength}}", hardMaxLength, -1)
	prompt = strings.Replace(prompt, "{{instructions}}", instructions, -1)
	prompt = strings.Replace(prompt, "{{commitMessageTemplate}}", commitMessageTemplate, -1)
	prompt = strings.Replace(prompt, "{{existingCommitMsg}}", existingCommitMsg, -1)

	completion, err := getCompletion(prompt)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(completion)
}
