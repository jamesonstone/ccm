package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	prompt "github.com/c-bata/go-prompt"
)

// promptMapping is a mapping of available conventional commits with emoji commits
func promptMapping(in prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "wip", Description: "work-in-progress 🚧 "},
		{Text: "articles", Description: "Store the article text posted by user"},
		{Text: "comments", Description: "Store the text commented to articles"},
		{Text: "groups", Description: "Combine users with specific rules"},
	}
	return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}

func main() {
	in := prompt.Input("ccm>> ", promptMapping,
		prompt.OptionTitle("CCM"),
		prompt.OptionPrefixTextColor(prompt.Yellow),
		prompt.OptionPreviewSuggestionTextColor(prompt.Blue),
		prompt.OptionSelectedSuggestionBGColor(prompt.LightGray),
		prompt.OptionSuggestionBGColor(prompt.DarkGray))

	if in == "wip" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Issue: ")
		issue, _ := reader.ReadString('\n')
		fmt.Print("Your commit: ")
		text, _ := reader.ReadString('\n')
		commit := "\"wip(" + strings.TrimSpace(issue) + "): 🚧 " + strings.TrimSpace(text) + "\""
		fmt.Println(commit)
		stdout, err := exec.Command("git commit -m " + commit).Output()

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Print(string(stdout))
	}
}
