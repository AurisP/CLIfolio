package commands

const careerText = `
## 💼 Career Overview

| Section        | Details                             |
|----------------|-------------------------------------|
| **Location**   | Tallinn, Estonia                    |
| **Education**  | Bachelor's, TalTech University      |
| **Position**   | Junior Software Engineer            |
| **Field**      | Embedded Linux Systems              |
| **Languages**  | C, Rust, Java                       |
| **Interests**  | Endurance sports, electronics       |

## 🛠 Skills & Technologies

| Skill     | Proficiency  | Description                        |
|-----------|--------------|------------------------------------|
| **C**     | Intermediate | Low-level embedded systems         |
| **Rust**  | Beginner     | Safe concurrency exploration       |
| **Java**  | Intermediate | Backend development                |
| **Git**   | Advanced     | Daily version control usage        |
| **Docker**| Intermediate | Containerized deployments          |

## 📁 Key Projects

| Project         | Description                              |
|-----------------|------------------------------------------|
| **FPV Drone**   | Control software, sensor integration     |
| **LED Lighting**| High-power LEDs, custom lighting effects |

## 🏃 Hobbies & Highlights

| Hobby                | Achievement                        |
|----------------------|------------------------------------|
| **Endurance Sports** | Full triathlon, ultra marathons    |
| **Electronics**      | Built custom LEDs, FPV drones      |
`

func Career() string {
	return careerText
}
