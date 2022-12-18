package utils

import "github.com/fatih/color"

func ShowLogo() {
	logo := "___________              .__.__    _________                  .___            \n\\_   _____/ _____ _____  |__|  |  /   _____/ ____   ____    __| _/___________ \n |    __)_ /     \\\\__  \\ |  |  |  \\_____  \\_/ __ \\ /    \\  / __ |/ __ \\_  __ \\\n |        \\  Y Y  \\/ __ \\|  |  |__/        \\  ___/|   |  \\/ /_/ \\  ___/|  | \\/\n/_______  /__|_|  (____  /__|____/_______  /\\___  >___|  /\\____ |\\___  >__|   \n        \\/      \\/     \\/                \\/     \\/     \\/      \\/    \\/       "
	println(logo + "\n\n\n")
	c := color.New(color.FgHiBlue, color.Bold)
	_, err := c.Println("R U N N I N G\n\n\n")
	if err != nil {
		return
	}
}
