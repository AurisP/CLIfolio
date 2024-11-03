package commands

const helpText = `
## 🛠 Available Commands

| Command  | Description                       |
|----------|-----------------------------------|
| **help** | Show available commands           |
| **about**| Learn a bit about me              |
| **career**| Take a look at my professional experience |
| **hobbies** | Check out what I do for fun     |
| **clear** | Clear the screen                 |
| **exit**  | Exit the application             |
`

func Help() string {
	return helpText
}
