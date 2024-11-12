package commands

const helpText = `
### Available Commands

- **help**: Show available commands
- **about**: Learn a bit about me
- **career**: Take a look at my professional experience
- **projects**: What I've been up to in my spare time
- **contacts**: Find ways to connect with me
- **clear**: Clear the screen
- **exit**: Exit the application

Use "help" anytime to rediscover commands and navigate through my CLI.

`

func Help() string {
	return helpText
}
