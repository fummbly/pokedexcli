package main

import "os"

// exit program
func commandExit(cfg *config, args ...string) error {
	os.Exit(0)
	return nil
}
