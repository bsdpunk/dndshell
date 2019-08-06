package commands

type Command struct {
	// Command Name
	Name string
	// Shorter maybe more common name
	ShortName string
	// Brief help entry for command
	Usage string
	// Launced Function on command
	Action func()
	// Action With String Argument
	StringAction func(s string)
	// Category of command
	Category string
	// List of child commands
	SubCommands Commands
}

type CommandsByName []Command

func (c CommandsByName) Len() int {
	return len(c)
}

func (c CommandsByName) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

// FullName returns the full name of the command.
// For subcommands this ensures that parent commands are part of the command path

type Commands []Command

func (c Command) Names() []string {
	names := []string{c.Name}

	if c.ShortName != "" {
		names = append(names, c.ShortName)
	}

	return names
	//return append(names, c.Aliases...)
}
func (cs Commands) HasCommand(name string) bool {
	for _, i := range cs {
		for _, n := range i.Names() {
			if n == name {
				return true
			}
		}
	}
	return false
}

func (c Command) HasName(name string) bool {
	for _, n := range c.Names() {
		if n == name {
			return true
		}
	}
	return false
}
func (cs Commands) NameIs(name string) Command {
	for _, c := range cs {
		if c.ShortName == name || c.Name == name {
			return c
		}
	}
	c := Command{}
	return c
}
