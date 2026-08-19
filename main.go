package main

import (
	"fmt"
	"os"
)

// clipboard_monitor - Monitor clipboard
func clipboard_monitor(path string) {
	fmt.Println("========================================")
	fmt.Println("  Clipboard-Monitor")
	fmt.Println("  Monitor clipboard")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	clipboard_monitor(path)
}
