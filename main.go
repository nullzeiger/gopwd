package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	fh "github.com/nullzeiger/gopwd/internal/filehandling"
	ph "github.com/nullzeiger/gopwd/internal/pwdhandling"
)

const FileName = ".pwds.csv"
const ExpectedArgs = 2

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

func main() {
	// Ensure the file exists
	if err := fh.Create(FileName); err != nil {
		log.Fatalf("Error creating file: %v", err)
	}

	if len(os.Args) < ExpectedArgs {
		log.Fatal("Expected 'all', 'create', 'delete', or 'search' command")
	}

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

func handleDelete() {
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

func handleSearch() {
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

func handleCreate(args []string) {
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
