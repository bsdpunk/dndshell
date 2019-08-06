package shell

import (
	"./character"
	//	class "./character/class"
	"./commands"
	"./dice"
	"./general"
	"fmt"
	"github.com/gobs/readline"
	"strings"
)

var found string = "no"
var list []string

var matches = make([]string, 0, len(list))

var coms = commands.Commands{
	{
		Name:      "Quit",
		ShortName: "quit",
		Usage:     "Exit the shell",
		Action:    general.End,
		Category:  "general",
	},
	{
		Name:      "Clear",
		ShortName: "clear",
		Usage:     "Clear the screen",
		Action:    general.Clear,
		Category:  "general",
	},
	{
		Name:         "Dice",
		ShortName:    "d",
		Usage:        "3d6: rolls 3 6 sided dice ",
		SubCommands:  dice.DiceSubs,
		StringAction: dice.GetDice,
		Category:     "general",
	},
	{
		Name:      "CreateCharacter",
		ShortName: "ic",
		Usage:     "",
		Action:    character.InteractiveCreateCharacter,
		Category:  "general",
	},
}

func AttemptedCompletion(text string, start, end int) []string {
	if start == 0 {
		return readline.CompletionMatches(text, CompletionEntry)
	} else {
		return nil
	}
}

func CompletionEntry(prefix string, index int) string {
	if index == 0 {
		matches = matches[:0]

		for _, w := range list {
			if strings.HasPrefix(w, prefix) {
				matches = append(matches, w)
			}
		}
	}

	if index < len(matches) {
		return matches[index]
	} else {
		return ""
	}
}
func NoAction() {
	fmt.Println("No action supplied with command")

}
func Shell() {

	for _, c := range coms {
		list = append(list, c.Name)
		list = append(list, c.ShortName)
	}
	//	cs := character.LoadClasses()
	prompt := "> "
	matches = make([]string, 0, len(list))
L:
	for {

		found = "no"
		readline.SetCompletionEntryFunction(CompletionEntry)
		readline.SetAttemptedCompletionFunction(nil)
		result := readline.ReadLine(&prompt)
		if result == nil {
			break L
		}

		input := *result
		words := strings.Fields(input)
		if len(words) > 0 && coms.HasCommand(words[0]) {
			cmd := coms.NameIs(words[0])
			if len(words) == 1 {
				if len(cmd.SubCommands) == 0 {
					cmd.Action()
				}
			} else {
				if cmd.SubCommands.HasCommand(words[1]) {
					cmd.Action()
				} else {
					cmd.StringAction(words[1])
				}
			}
		}
	}

	return
}
