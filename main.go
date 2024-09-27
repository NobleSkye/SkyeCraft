package main

import (
	"os"
	"path/filepath"

	webview "github.com/webview/webview_go"
)

func main() {
	// Create a new webview instance
	w := webview.New(true) // `true` enables debug mode
	defer w.Destroy()

	// Set window title and dimensions
	w.SetTitle("SkyeCraft")
	w.SetSize(800, 600, webview.HintNone)

	// Get the current working directory
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// Build the full path to the HTML file
	htmlPath := filepath.Join(dir, "skyecraft.html")

	// Load the local HTML file
	w.Navigate("file://" + htmlPath)

	// Run the webview instance
	w.Run()
}
