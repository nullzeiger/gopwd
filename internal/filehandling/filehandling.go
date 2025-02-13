// Copyright 2025 Ivan Guerreschi <ivan.guerreschi.dev@gmail.com>.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package containing file handling functions
package filehandling

import "os"

const perm = 0o644

// Create creates a file otherwise an error
func Create(name string) error {
	name = fileName(name)

	exist := fileExists(name)
	if exist {
		return nil
	}

	file, err := os.Create(name)
	if err != nil {
		return err
	}

	defer file.Close()

	return nil
}

// Open opens a file and returns File for subsequent I/O operations or an error
func Open(name string) (*os.File, error) {
	name = fileName(name)

	file, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_RDWR, perm)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// Remove removes a file otherwise returns an error
func Remove(name string) error {
	name = fileName(name)

	err := os.Remove(name)
	if err != nil {
		return err
	}

	return nil
}

// fileExists returns a boolean true if the file exists
// otherwise false if it does not exist
func fileExists(name string) bool {
	_, err := os.Stat(name)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	return false
}

// getHome returns a string with the value of the home directory
// otherwise a panic
func getHome() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	return home
}

// fileName returns a string with the concatenated value
// of the home directory and the file name
func fileName(name string) string {
	name = getHome() + "/" + name

	return name
}
