package commands

const helpText = `
Available Commands:

  help      Show available commands
  about     Learn a bit about me
  career    Take a look at my professional experience
  contacts  Where to find me
  clear     Clear the screen
  exit      Exit the application
`

func Help() string {
	return helpText
}
