package main

import (
	"testing"
)

func TestGetMessage(t *testing.T) {
	expected := "Hello World!"
    actual := GetMessage()
    if actual != expected {
       t.Errorf("GetMessage was incorrect, got: %v, want: %v.", actual, expected)
    }
}
