// Copyright 2025 Ivan Guerreschi <ivan.guerreschi.dev@gmail.com>.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Test filehandling package
package filehandling

import "testing"

const filename = "gopwdtestfile"

// TestCreate test Create() function
func TestCreate(t *testing.T) {
	err := Create(filename)
	if err != nil {
		t.Errorf("Create() error %v", err)
	}
}

// TestOpen test Open() function
func TestOpen(t *testing.T) {
	_, err := Open(filename)
	if err != nil {
		t.Errorf("Open() error %v", err)
	}
}

// TestRemove test Remove() function
func TestRemove(t *testing.T) {
	err := Remove(filename)
	if err != nil {
		t.Errorf("Remove() error %v", err)
	}
}
