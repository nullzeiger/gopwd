// Copyright 2025 Ivan Guerreschi
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package commands contains handle functions for user input
package commands

import (
	"flag"
	"fmt"
	"log"
	"os"

	fh "github.com/nullzeiger/gopwd/internal/filehandling"
	ph "github.com/nullzeiger/gopwd/internal/pwdhandling"
)

// FileName file for storage password
const FileName = ".pwds.csv"

func printSlice(s []string) {
	for _, pwd := range s {
		fmt.Println(pwd)
	}
}

// withFile is a helper to open a file and pass it to a function.
func withFile(fn func(*os.File) error) error {
	file, err := fh.Open(FileName)
	if err != nil {
		return err
	}
	defer file.Close()

	return fn(file)
}

// HandleAll print all password
func HandleAll() {
	if err := withFile(func(file *os.File) error {
		pwds, err := ph.All(file)
		if err != nil {
			return err
		}
		printSlice(pwds)
		return nil
	}); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

// HandleDelete delete password
func HandleDelete() {
	deleteFlag := flag.NewFlagSet("delete", flag.ExitOnError)
	index := deleteFlag.Int("index", -1, "Index of password to delete")

	err := deleteFlag.Parse(os.Args[2:])
	if err != nil {
		log.Fatalf("Error parsing flags: %v", err)
	}

	if *index < 0 {
		log.Fatal("Please provide a valid index (>= 0)")
	}

	if err := withFile(func(file *os.File) error {
		ok, err := ph.Delete(file, *index)
		if err != nil {
			return err
		}
		fmt.Println("Deleted:", ok)
		return nil
	}); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

// HandleSearch search password
func HandleSearch() {
	searchFlag := flag.NewFlagSet("search", flag.ExitOnError)
	query := searchFlag.String("query", "", "Search query")

	err := searchFlag.Parse(os.Args[2:])
	if err != nil {
		log.Fatalf("Error parsing flags: %v", err)
	}

	if *query == "" {
		log.Fatal("Please provide a search query using --query")
	}

	if err := withFile(func(file *os.File) error {
		pwds, err := ph.Search(file, *query)
		if err != nil {
			return err
		}
		printSlice(pwds)
		return nil
	}); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

// HandleCreate create password
func HandleCreate(args []string) {
	createFlag := flag.NewFlagSet("create", flag.ExitOnError)
	name := createFlag.String("name", "", "Name")
	username := createFlag.String("username", "", "Username")
	email := createFlag.String("email", "", "Email")
	password := createFlag.String("password", "", "Password")

	err := createFlag.Parse(args)
	if err != nil {
		log.Fatalf("Error parsing flags: %v", err)
	}

	if *name == "" || *username == "" || *email == "" || *password == "" {
		log.Fatal("All fields --name, --username, --email, and --password are required")
	}

	if err := withFile(func(file *os.File) error {
		pwd := ph.Pwd{
			Name:     *name,
			Username: *username,
			Email:    *email,
			Password: *password,
		}

		if err := ph.Create(file, pwd); err != nil {
			return err
		}
		fmt.Println("Password created successfully")
		return nil
	}); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
