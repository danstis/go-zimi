package main

import (
	"log"

	"github.com/danstis/go-zimi/internal/version"
)

// Main entry point for the app.
func main() {
	log.Printf("Version %q", version.Version)
}
