// Copyright 2025 Ivan Guerreschi <ivan.guerreschi.dev@gmail.com>
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package filehandling contains utility functions for managing files.
package filehandling

import (
	"fmt"
	"os"
	"path/filepath"
)

const filePerm = 0o644

// Create ensures the file exists in the user's home directory.
// If it does not exist, it creates it.
func Create(name string) error {
	fullPath := fullPath(name)

	if fileExists(fullPath) {
		return nil
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	return nil
}

// Open opens a file in append+read-write mode and creates it if it doesn't exist.
func Open(name string) (*os.File, error) {
	fullPath := fullPath(name)

	file, err := os.OpenFile(fullPath, os.O_APPEND|os.O_CREATE|os.O_RDWR, filePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	return file, nil
}

// Remove deletes the file from the home directory.
func Remove(name string) error {
	fullPath := fullPath(name)

	if err := os.Remove(fullPath); err != nil {
		return fmt.Errorf("failed to remove file: %w", err)
	}

	return nil
}

// fileExists returns true if the file exists, false otherwise.
func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// getHomeDir returns the user's home directory or panics if it cannot be determined.
func getHomeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic("unable to determine user home directory: " + err.Error())
	}
	return home
}

// fullPath returns the absolute path of the file in the user's home directory.
func fullPath(name string) string {
	return filepath.Join(getHomeDir(), name)
}

