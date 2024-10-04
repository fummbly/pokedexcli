package main

import "os"

// exit program
func commandExit(cfg *config) error {
	os.Exit(0)
	return nil
}
