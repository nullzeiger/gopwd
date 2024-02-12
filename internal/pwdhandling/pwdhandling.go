// Copyright 2024 Ivan Guerreschi <ivan.guerreschi.dev@gmail.com>.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package containing password handling functions
package pwdhandling

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

// Structure that defines the fields for Pwd
type Pwd struct {
	Name     string
	Username string
	Email    string
	Password string
}

// All returns a slice of strings containing all the passwords
func All(file *os.File) []string {
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	pwds := make([]string, 0, len(records))

	for i, item := range records {
		pwd := Pwd{
			Name:     item[0],
			Username: item[1],
			Email:    item[2],
			Password: item[3],
		}

		pwds = append(pwds,
			fmt.Sprintf("[%d] Name: %s Username: %s Email: %s Password: %s\n",
				i, pwd.Name, pwd.Username, pwd.Email, pwd.Password))
	}

	return pwds
}

// Delete() deletes a line indicated by the key function parameter and returns true if successful
func Delete(file *os.File, key int) bool {
	defer file.Close()

	const tempFile = ".pwdTmp.csv"

	fileTmp, err := os.CreateTemp("", tempFile)
	if err != nil {
		panic(err)
	}

	defer os.Remove(file.Name())

	scanner := bufio.NewScanner(file)

	i := 1

	for scanner.Scan() {
		line := scanner.Text()
		if i != key {
			_, err = fileTmp.WriteString(line + "\n")
			if err != nil {
				panic(err)
			}
		}
		i++
	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	}

	err = os.Rename(tempFile, file.Name())
	if err != nil {
		panic(err)
	}

	return true
}

// Search returns a slice with all the passwords found with the word used for the search
func Search(file *os.File, key string) []string {
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	pwds := make([]string, 0, len(records))

	for i, item := range records {
		if key == item[0] || key == item[1] || key == item[2] || key == item[3] {
			pwds = append(pwds,
				fmt.Sprintf("[%d] Name: %s Username: %s Email: %s Password: %s\n",
					i, item[0], item[1], item[2], item[3]))
		}
	}

	return pwds
}

// Write writes a new password to the file otherwise an error
func Write(file *os.File, pwd Pwd) error {
	defer file.Close()

	newPassword := fmt.Sprintf("%s,%s,%s,%s\n", pwd.Name, pwd.Username, pwd.Email, pwd.Password)

	_, err := file.WriteString(newPassword)
	if err != nil {
		panic(err)
	}

	return nil
}
