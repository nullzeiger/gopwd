// Copyright 2022 Ivan Guerreschi <ivan.guerreschi.dev@gmail.com>.
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

// Print slice
func printSlice(s []string) {
	for _, pwd := range s {
		fmt.Printf("%v\n", pwd)
	}
}

func main() {
	// Create file is not exist
	err := fh.Create(fileName)
	if err != nil {
		log.Fatal("Error create file", err)
	}

	// Flags
	allFlag := flag.Bool("all", false, "display all password")
	deleteFlag := flag.Int("delete", 0, "delete password")
	searchFlag := flag.String("search", "", "search password")
	// New subcommand
	createSub := flag.NewFlagSet("create", flag.ExitOnError)
	// Subcommad flags
	name := createSub.String("name", "", "name")
	username := createSub.String("username", "", "username")
	email := createSub.String("email", "", "email")
	password := createSub.String("password", "", "password")

	flag.Parse()

	if len(os.Args) <= 1 {
		log.Fatal("expected 'all', 'create', 'delete', 'search' commands")
	}

	if *allFlag {
		file, err := fh.Open(fileName)
		if err != nil {
			log.Fatal("Error open file", err)
		}

		pwds, err := ph.All(file)
		if err != nil {
			log.Fatal("Error read file", err)
		}

		printSlice(pwds)

		file.Close()

		return
	}

	if *deleteFlag > 0 {
		file, err := fh.Open(fileName)
		if err != nil {
			log.Fatal("Error open file", err)
		}

		pwds, err := ph.Delete(file, *deleteFlag)
		if err != nil {
			log.Fatal("Error delete file", err)
		}

		fmt.Println(pwds)

		file.Close()

		return
	}

	if *searchFlag != "" {
		file, err := fh.Open(fileName)
		if err != nil {
			log.Fatal("Error open file", err)
		}

		pwds, err := ph.Search(file, *searchFlag)
		if err != nil {
			log.Fatal("Error search file", err)
		}

		for _, pwd := range pwds {
			fmt.Println(pwd)
		}

		file.Close()

		return
	}

	if os.Args[1] == "create" {
		if len(os.Args[2:]) == 0 {
			log.Fatal("expected 'name', 'username', 'email', 'password' arguments")
		}

		file, err := fh.Open(fileName)
		if err != nil {
			log.Fatal("Error open file", err)
		}

		err = createSub.Parse(os.Args[2:])
		if err != nil {
			log.Fatal("Error parse subcommands")
		}

		pwd := ph.Pwd{
			Name:     *name,
			Username: *username,
			Email:    *email,
			Password: *password,
		}

		n, err := ph.Create(file, pwd)
		if err != nil {
			log.Fatal("Erorr create password", err)
		}

		if n > 0 {
			fmt.Println("Pasword create")
		}

		file.Close()

		return
	}
}
