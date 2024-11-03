package commands

const careerText = `
Here's a list of my professional experiences:

Artec Design — Software Engineer
June 2024 - Present
- Developing applications for Linux and Cortex-A systems.
- Working on embedded systems with Cortex-M processors.

Adacel, TalTech — Project Team Member
September 2023 - March 2024
- Developed a video latency measuring system for remote air traffic control (ATC).
- Collaborative project with Adacel Technologies and TalTech University.

TalTech — Maze Algorithms Project Lead
January 2022 - June 2022
- Led a project to develop a visualization program for maze algorithms.
- Implemented maze generation algorithms and managed a small team.

VokBikes — Electronics Assembler
June 2022 - September 2022
- Assembled electronics for electric cargo bikes.
- Diagnosed and repaired electrical issues in malfunctioning bikes.   |
`

func Career() string {
	return careerText
}
