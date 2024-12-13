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
	Name, Username, Email, Password string
}

// All returns a slice of strings containing all the passwords
func All(file *os.File) ([]string, error) {
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
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
			fmt.Sprintf("[%d] Name: %s Username: %s Email: %s Password: %s",
				i, pwd.Name, pwd.Username, pwd.Email, pwd.Password))
	}

	return pwds, nil
}

// Create writes a new password to the file otherwise an error
func Create(file *os.File, pwd Pwd) (int, error) {
	defer file.Close()

	newPassword := fmt.Sprintf("%s,%s,%s,%s\n", pwd.Name, pwd.Username, pwd.Email, pwd.Password)

	n, err := file.WriteString(newPassword)
	if err != nil {
		return 0, err
	}

	return n, nil
}

// Delete() deletes a line indicated by the key function parameter and returns true if successful
func Delete(file *os.File, key int) (bool, error) {
	defer file.Close()

	fileTemp, err := os.CreateTemp("", "pwdTemp.csv")
	if err != nil {
		return false, err
	}

	defer fileTemp.Close()
	defer os.Remove(fileTemp.Name())

	scanner := bufio.NewScanner(file)

	i := 1

	for scanner.Scan() {
		line := scanner.Text()
		if i != (key - 1) {
			_, err = fileTemp.WriteString(line + "\n")
			if err != nil {
				return false, err
			}
		}

		i++
	}

	err = scanner.Err()
	if err != nil {
		return false, err
	}

	err = os.Rename(fileTemp.Name(), file.Name())
	if err != nil {
		return false, err
	}

	return true, nil
}

// Search returns a slice with all the passwords found with the word used for the search
func Search(file *os.File, key string) ([]string, error) {
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	pwds := make([]string, 0, len(records))

	for i, item := range records {
		if key == item[0] || key == item[1] || key == item[2] || key == item[3] {
			pwds = append(pwds,
				fmt.Sprintf("[%d] Name: %s Username: %s Email: %s Password: %s",
					i, item[0], item[1], item[2], item[3]))
		}
	}

	return pwds, nil
}
