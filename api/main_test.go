package main

import (
	"os"
	"testing"
)


func TestMain(m *testing.M) {
	code := m.Run()
	PrintGreen("[OK] main test successfull")
	os.Exit(code)
}
