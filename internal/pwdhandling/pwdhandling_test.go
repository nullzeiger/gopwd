// Copyright 2024 Ivan Guerreschi <ivan.guerreschi.dev@gmail.com>.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Test pwdhandling package
package pwdhandling

import (
	"os"
	"testing"
)

// TestAll test All() function
func TestAll(t *testing.T) {
	file, err := os.CreateTemp("", "pwdhandlingtest.csv")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	defer os.Remove(file.Name())

	pwd := Pwd{"Google", "mr", "mario.rossi@google.com", "1234"}

	_, err = Create(file, pwd)
	if err != nil {
		panic(err)
	}

	file, err = os.Open(file.Name())
	if err != nil {
		panic(err)
	}

	want := 1

	got, err := All(file)
	if err != nil || len(got) == 0 {
		t.Errorf("All() error %v got %d wanted %d", err, len(got), want)
	}
}

// TestCreate test Create() function
func TestCreate(t *testing.T) {
	file, err := os.CreateTemp("", "pwdhandlingtest.csv")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	defer os.Remove(file.Name())

	pwd := Pwd{"Google", "mr", "mario.rossi@google.com", "1234"}

	want := 1

	got, err := Create(file, pwd)
	if err != nil || got <= want {
		t.Errorf("Create() error %v got %x wanted %d", err, got, want)
	}
}

// TestDelete test Delete() function
func TestDelete(t *testing.T) {
	file, err := os.CreateTemp("", "pwdhandlingtest.csv")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	defer os.Remove(file.Name())

	pwd1 := Pwd{"Google", "ff", "Fred.Flintstone.@google.com", "1234"}

	_, err = Create(file, pwd1)
	if err != nil {
		panic(err)
	}

	file, err = os.Open(file.Name())
	if err != nil {
		panic(err)
	}

	want := true

	got, err := Delete(file, 1)
	if err != nil || got != want {
		t.Errorf("Search() error %v got %t wanted %t", err, got, want)
	}
}

// TestSearch test Search() function
func TestSearch(t *testing.T) {
	file, err := os.CreateTemp("", "pwdhandlingtest.csv")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	defer os.Remove(file.Name())

	pwd1 := Pwd{"Google", "ff", "Fred.Flintstone.@google.com", "1234"}

	_, err = Create(file, pwd1)
	if err != nil {
		panic(err)
	}

	file, err = os.Open(file.Name())
	if err != nil {
		panic(err)
	}

	want := 1

	got, err := Search(file, "ff")
	if err != nil || len(got) < want {
		t.Errorf("Search() error %v got %d wanted %d", err, len(got), want)
	}
}
