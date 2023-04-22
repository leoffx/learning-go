package main

import (
	"testing"
)

func TestCheckWinTrue(t *testing.T) {
	initializeBoard()
	makeMove("X", 1, 1)
	makeMove("X", 1, 2)
	makeMove("X", 1, 3)
	if !checkWin() {
		t.Fatalf("Expected checkWin to return true, but it returned false")
	}
}

func TestCheckWinFalse(t *testing.T) {
	initializeBoard()
	makeMove("X", 1, 1)
	makeMove("X", 1, 2)
	makeMove("O", 1, 3)
	if checkWin() {
		t.Fatalf("Expected checkWin to return false, but it returned true")
	}
}
