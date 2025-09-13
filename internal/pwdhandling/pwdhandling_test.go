// Copyright 2025 Ivan Guerreschi
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Test pwdhandling package
package pwdhandling

import (
	"os"
	"testing"
)

// helper per creare un file temporaneo e rimuoverlo al termine
func tempFile(t *testing.T) *os.File {
	t.Helper()

	file, err := os.CreateTemp("", "pwdhandlingtest.csv")
	if err != nil {
		t.Fatalf("cannot create temp file: %v", err)
	}

	t.Cleanup(func() {
		file.Close()
		os.Remove(file.Name())
	})

	return file
}

// TestAll tests All() function
func TestAll(t *testing.T) {
	file := tempFile(t)

	pwd := Pwd{"Google", "mr", "mario.rossi@google.com", "1234"}
	if err := Create(file, pwd); err != nil {
		t.Fatalf("Create() failed: %v", err)
	}

	// Ri-apro il file in lettura
	file.Close()

	file, err := os.Open(file.Name())
	if err != nil {
		t.Fatalf("cannot reopen file: %v", err)
	}
	defer file.Close()

	got, err := All(file)
	if err != nil {
		t.Fatalf("All() error: %v", err)
	}

	if len(got) != 1 {
		t.Errorf("All() got %d, want %d", len(got), 1)
	}
}

// TestCreate tests Create() function
func TestCreate(t *testing.T) {
	file := tempFile(t)

	pwd := Pwd{"Google", "mr", "mario.rossi@google.com", "1234"}

	if err := Create(file, pwd); err != nil {
		t.Fatalf("Create() error: %v", err)
	}

	// Check that it was really written
	file.Close()

	file, err := os.Open(file.Name())
	if err != nil {
		t.Fatalf("cannot reopen file: %v", err)
	}
	defer file.Close()

	records, err := All(file)
	if err != nil {
		t.Fatalf("All() error: %v", err)
	}

	if len(records) != 1 {
		t.Errorf("Create() did not write record, got %d, want %d", len(records), 1)
	}
}

// TestDelete tests Delete() function
func TestDelete(t *testing.T) {
	file := tempFile(t)

	pwd := Pwd{"Google", "ff", "Fred.Flintstone.@google.com", "1234"}

	if err := Create(file, pwd); err != nil {
		t.Fatalf("Create() failed: %v", err)
	}

	// Reopen for delete
	file.Close()

	file, err := os.Open(file.Name())
	if err != nil {
		t.Fatalf("cannot reopen file: %v", err)
	}
	defer file.Close()

	got, err := Delete(file, 0)
	if err != nil {
		t.Fatalf("Delete() error: %v", err)
	}

	if !got {
		t.Errorf("Delete() got %t, want %t", got, true)
	}
}

// TestSearch tests Search() function
func TestSearch(t *testing.T) {
	file := tempFile(t)

	pwd := Pwd{"Google", "ff", "Fred.Flintstone.@google.com", "1234"}
	if err := Create(file, pwd); err != nil {
		t.Fatalf("Create() failed: %v", err)
	}

	// Reopen for search
	file.Close()

	file, err := os.Open(file.Name())
	if err != nil {
		t.Fatalf("cannot reopen file: %v", err)
	}

	defer file.Close()

	got, err := Search(file, "ff")
	if err != nil {
		t.Fatalf("Search() error: %v", err)
	}

	if len(got) < 1 {
		t.Errorf("Search() got %d, want at least %d", len(got), 1)
	}
}
