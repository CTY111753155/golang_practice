package main

import "testing"

// TestRun is disabled by default because run() requires
// command-line flags and a live database connection.
// Enable manually if you really want to integration-test startup.
func TestRun(t *testing.T) {
	t.Skip("Skipping TestRun: requires DB and flags")
}
