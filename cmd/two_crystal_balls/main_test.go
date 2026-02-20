package main

import "testing"

func TestBreakFound(t *testing.T) {
	breaks := []bool{false, false, false, false, false, false, true, true, true}
	if twoCrystalBalls(breaks) != 6 {
		t.Errorf("Expected 6")
	}
}

func TestBreakFoundTwo(t *testing.T) {
	breaks := []bool{
		false, false, false, false, false, false, false, false, false, false, true, true, true,
	}

	if twoCrystalBalls(breaks) != 10 {
		t.Errorf("Expected 10")
	}
}

func TestBreakNotFound(t *testing.T) {
	breaks := []bool{false, false, false, false, false, false, false, false, false}

	if twoCrystalBalls(breaks) != -1 {
		t.Errorf("Expected -1")
	}
}

func TestBreakNotFoundTwo(t *testing.T) {
	breaks := []bool{
		false, false, false, false, false, false, false, false, false, false, false, false,
		false,
	}

	if twoCrystalBalls(breaks) != -1 {
		t.Errorf("Expected -1")
	}
}
