package commands

const projectsText = `
### 🔧 Personal Projects

Over the years, I've worked on various personal projects, often incorporating electronics for a hands-on experience.  
Here are a few highlights of my individual projects outside of work and studies:

1. **High-Powered LED Flashlight**  
   Built a custom **140W battery-powered LED flashlight** with active cooling, repurposing an old CPU cooler to manage heat output.

2. **FPV Drone Enthusiast**  
   Spent 2 years in the **FPV drone community**, building, maintaining, and flying my own custom FPV drone for immersive flight experiences.

3. **WiFi-Controlled RC Car**  
   Modified an RC car with an **ESP32 controller**, enabling wireless control over WiFi for remote navigation and experiments in IoT.

These projects have deepened my technical skills and allowed me to combine software, hardware, and creativity in practical ways!
`

func Projects() string {
	return projectsText
}
