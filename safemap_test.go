// Copyright 2015 hit9. All rights reserved.

package safemap

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

// Assert util.
func assert(t *testing.T, b bool) {
	if !b {
		_, fileName, line, _ := runtime.Caller(1)
		cwd, err := os.Getwd()
		if err != nil {
			t.Errorf("unexcepted:%v", err)
		}
		fileName, err = filepath.Rel(cwd, fileName)
		if err != nil {
			t.Errorf("unexcepted:%v", err)
		}
		t.Errorf("\nassertion failed: %s:%d", fileName, line)
	}
}

func TestBasic(t *testing.T) {
	m := New()
	// Set
	m.Set("key", "val")
	assert(t, m.Len() == 1)
	// Get
	val, ok := m.Get("key")
	assert(t, ok)
	assert(t, val == "val")
	// Items
	d := map[string]string{"key": "val"}
	assert(t, len(d) == m.Len())
	assert(t, m.Items()["key"] == "val")
	// Len
	assert(t, m.Len() == 1)
	// Delete
	assert(t, m.Delete("key"))
	assert(t, !m.Delete("key1"))
	assert(t, m.Len() == 0)
	// Clear
	m.Set("key", "val")
	m.Clear()
	assert(t, m.Len() == 0)
}
