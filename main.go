// Copyright 2025 Ivan Guerreschi <ivan.guerreschi.dev@gmail.com>.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	fh "github.com/nullzeiger/gopwd/internal/filehandling"
	ph "github.com/nullzeiger/gopwd/internal/pwdhandling"
)

const fileName = ".pwds.csv"

func printSlice(s []string) {
	for _, pwd := range s {
		fmt.Println(pwd)
	}
}

func main() {
	// Ensure the file exists
	if err := fh.Create(fileName); err != nil {
		log.Fatalf("Error creating file: %v", err)
	}

	if len(os.Args) < 2 {
		log.Fatal("Expected 'all', 'create', 'delete', or 'search' command")
	}

	// Subcommand parsing
	switch os.Args[1] {
	case "all":
		handleAll()
	case "delete":
		handleDelete()
	case "search":
		handleSearch()
	case "create":
		handleCreate(os.Args[2:])
	default:
		log.Fatalf("Unknown command: %s", os.Args[1])
	}
}

func handleAll() {
	file, err := fh.Open(fileName)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	pwds, err := ph.All(file)
	if err != nil {
		log.Fatalf("Error reading passwords: %v", err)
	}

	printSlice(pwds)
}

func handleDelete() {
	deleteFlag := flag.NewFlagSet("delete", flag.ExitOnError)
	index := deleteFlag.Int("index", 0, "Index of password to delete")
	deleteFlag.Parse(os.Args[2:])

	if *index <= 0 {
		log.Fatal("Please provide a valid index to delete (greater than 0)")
	}

	file, err := fh.Open(fileName)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	result, err := ph.Delete(file, *index)
	if err != nil {
		log.Fatalf("Error deleting password: %v", err)
	}

	fmt.Println(result)
}

func handleSearch() {
	searchFlag := flag.NewFlagSet("search", flag.ExitOnError)
	query := searchFlag.String("query", "", "Search query")
	searchFlag.Parse(os.Args[2:])

	if *query == "" {
		log.Fatal("Please provide a search query using --query")
	}

	file, err := fh.Open(fileName)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	pwds, err := ph.Search(file, *query)
	if err != nil {
		log.Fatalf("Error searching passwords: %v", err)
	}

	printSlice(pwds)
}

func handleCreate(args []string) {
	createFlag := flag.NewFlagSet("create", flag.ExitOnError)
	name := createFlag.String("name", "", "Name")
	username := createFlag.String("username", "", "Username")
	email := createFlag.String("email", "", "Email")
	password := createFlag.String("password", "", "Password")
	createFlag.Parse(args)

	if *name == "" || *username == "" || *email == "" || *password == "" {
		log.Fatal("All fields --name, --username, --email, and --password are required")
	}

	file, err := fh.Open(fileName)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	pwd := ph.Pwd{
		Name:     *name,
		Username: *username,
		Email:    *email,
		Password: *password,
	}

	n, err := ph.Create(file, pwd)
	if err != nil {
		log.Fatalf("Error creating password: %v", err)
	}

	if n > 0 {
		fmt.Println("Password created successfully")
	}
}

