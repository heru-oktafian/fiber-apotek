package helpers

import (
	"fmt"
	"os"
)

func PrintFiberLikeBanner(vApp, vHost string, vPort int, handlers int) {
	pid := os.Getpid()

	banner := fmt.Sprintf(
		"\x1b[36m┌───────────────────────────────────────────────────┐\n"+
			"│ \x1b[1m\x1b[37m%-49s \x1b[36m│\n"+
			"│ \x1b[90m%-49s \x1b[36m│\n"+
			"│ \x1b[36m%-36s \x1b[36m│\n"+
			"│                                                   │\n"+
			"│ \x1b[33mHandlers ............ %-3d \x1b[34mProcesses ........... 1 \x1b[36m│\n"+
			"│ \x1b[32mPrefork ....... Disabled  \x1b[31mPID ............. %-5d \x1b[36m│\n"+
			"└───────────────────────────────────────────────────┘\x1b[0m\n",
		centerText(vApp, 49),
		centerText(fmt.Sprintf("http://%s:%d", vHost, vPort), 49),
		centerText(fmt.Sprintf("(bound on host %s and port %d)", vHost, vPort), 49),
		handlers,
		pid,
	)

	fmt.Print(banner)
}

func centerText(text string, width int) string {
	if len(text) >= width {
		return text
	}

	padding := width - len(text)
	left := padding / 2
	right := padding - left

	return repeat(" ", left) + text + repeat(" ", right)
}

func repeat(s string, count int) string {
	result := ""
	for i := 0; i < count; i++ {
		result += s
	}
	return result
}
