// Copyright 2025 Ivan Guerreschi
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package main contains parsing user input
package main

import (
	"log"
	"os"

	cmd "github.com/nullzeiger/gopwd/internal/commands"
	fh "github.com/nullzeiger/gopwd/internal/filehandling"
)

// ExpectedArgs number of arguments os.Args[0] and os.Args[1]
const ExpectedArgs = 2

func main() {
	// Ensure the file exists
	if err := fh.Create(cmd.FileName); err != nil {
		log.Fatalf("Error creating file: %v", err)
	}

	if len(os.Args) < ExpectedArgs {
		log.Fatal("Expected 'all', 'create', 'delete', or 'search' command")
	}

	switch os.Args[1] {
	case "all":
		cmd.HandleAll()
	case "delete":
		cmd.HandleDelete()
	case "search":
		cmd.HandleSearch()
	case "create":
		cmd.HandleCreate(os.Args[2:])
	default:
		log.Fatalf("Unknown command: %s", os.Args[1])
	}
}
