package main

import (
	"fmt"

	prompt "github.com/c-bata/go-prompt"
)

//
func completer(in prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "wip", Description: "work-in-progress 🚧 "},
		{Text: "articles", Description: "Store the article text posted by user"},
		{Text: "comments", Description: "Store the text commented to articles"},
		{Text: "groups", Description: "Combine users with specific rules"},
	}
	return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}

func main() {
	in := prompt.Input("ccm>> ", completer,
		prompt.OptionTitle("CCM"),
		prompt.OptionPrefixTextColor(prompt.Yellow),
		prompt.OptionPreviewSuggestionTextColor(prompt.Blue),
		prompt.OptionSelectedSuggestionBGColor(prompt.LightGray),
		prompt.OptionSuggestionBGColor(prompt.DarkGray))

	if in == "wip" {
		fmt.Println("wip(issue-123): 🚧 ")
	}
}
