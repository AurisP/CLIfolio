package commands

const aboutText = `
Hi, I’m Auris!

I'm a software developer currently based in Tallinn, Estonia.I graduated with honors from TalTech University specializing in Hardware and Software development.
My primary focus is on embedded Linux systems, where I work extensively with C, Rust, and Java.

Both at and outside of work I'm constantly exploring new technologies. 
For example, this terminal application was built with Go and Wish, technologies I had no previous experience with and wanted to learn out of curiosity.

Outside of software development, I’m passionate about ultra sports like full distance triathlons and ultra marathons. 
I like to push myself both physically and mentally and sport gives me another way to do it.

I love to do long distance hikes in very rural places to feel close to nature and experience something different from everyday life.
`

func About() string {

	return aboutText
}
