package commands

import (
	"log"

	"github.com/charmbracelet/glamour"
)

const aboutText = `# About

I'm a software developer with a passion for programming!

## What I Love
- Creating interactive command-line interfaces
- Using Go for high-performance applications
- Exploring new frameworks

For more information, check out my projects on [GitHub](https://github.com).`

func About() string {
	renderedText, err := glamour.Render(aboutText, "dark")
	if err != nil {
		log.Println("Error rendering with Glamour:", err)
		return "Error displaying content."
	}
	return renderedText
}
