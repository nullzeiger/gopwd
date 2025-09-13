// Copyright 2025 Ivan Guerreschi
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package pwdhandling contains password handling functions
package pwdhandling

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

// Pwd represents a single password entry.
type Pwd struct {
	Name, Username, Email, Password string
}

// formatPwd returns a human-readable string for a password entry.
func formatPwd(index int, p Pwd) string {
	return fmt.Sprintf("[%d] Name: %s Username: %s Email: %s Password: %s",
		index, p.Name, p.Username, p.Email, p.Password)
}

// All returns a slice of strings containing all the passwords.
func All(file *os.File) ([]string, error) {
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	pwds := make([]string, 0, len(records))
	for i, item := range records {
		if len(item) < 4 {
			continue // skip malformed rows
		}
		pwd := Pwd{item[0], item[1], item[2], item[3]}
		pwds = append(pwds, formatPwd(i, pwd))
	}

	return pwds, nil
}

// Create appends a new password to the file.
func Create(file *os.File, pwd Pwd) error {
	writer := csv.NewWriter(file)
	defer writer.Flush()

	record := []string{pwd.Name, pwd.Username, pwd.Email, pwd.Password}
	if err := writer.Write(record); err != nil {
		return err
	}
	return nil
}

// Delete removes a row by index and rewrites the file.
func Delete(file *os.File, key int) (bool, error) {
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return false, err
	}

	if key < 0 || key >= len(records) {
		return false, fmt.Errorf("invalid index %d", key)
	}

	// Remove the row
	records = append(records[:key], records[key+1:]...)

	// Write to a temporary file
	tempFile, err := os.CreateTemp("", "pwdTemp_*.csv")
	if err != nil {
		return false, err
	}
	defer os.Remove(tempFile.Name())

	writer := csv.NewWriter(tempFile)
	if err := writer.WriteAll(records); err != nil {
		tempFile.Close()
		return false, err
	}
	writer.Flush()
	tempFile.Close()

	// Replace the original file
	if err := os.Rename(tempFile.Name(), file.Name()); err != nil {
		return false, err
	}

	return true, nil
}

// Search looks for passwords matching the query (case-insensitive, substring match).
func Search(file *os.File, query string) ([]string, error) {
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var results []string
	q := strings.ToLower(query)

	for i, row := range records {
		if len(row) < 4 {
			continue
		}
		pwd := Pwd{row[0], row[1], row[2], row[3]}
		if strings.Contains(strings.ToLower(pwd.Name), q) ||
			strings.Contains(strings.ToLower(pwd.Username), q) ||
			strings.Contains(strings.ToLower(pwd.Email), q) ||
			strings.Contains(strings.ToLower(pwd.Password), q) {
			results = append(results, formatPwd(i, pwd))
		}
	}

	return results, nil
}
