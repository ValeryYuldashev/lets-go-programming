package internal

import (
	"testing"
)

func TestAddition(t *testing.T) {
	var result int
	var err error
	result, err = Calculate(1, 2, "+")
	if err != nil {
		t.Error("Error isn't nil")
	}
	if result != 3 {
		t.Error("Expected 3, got ", result)
	}
}

func TestSubstract(t *testing.T) {
	var result int
	var err error
	result, err = Calculate(1, 2, "-")
	if err != nil {
		t.Error("Error isn't nil")
	}
	if result != -1 {
		t.Error("Expected -1, got ", result)
	}
}

func TestMultiplication(t *testing.T) {
	var result int
	var err error
	result, err = Calculate(2, 2, "*")
	if err != nil {
		t.Error("Error isn't nil")
	}
	if result != 4 {
		t.Error("Expected 4, got ", result)
	}
}

func TestDivision(t *testing.T) {
	var result int
	var err error
	result, err = Calculate(2, 2, "/")
	if err != nil {
		t.Error("Error isn't nil")
	}
	if result != 1 {
		t.Error("Expected 1, got ", result)
	}
}

func TestDivisionByZero(t *testing.T) {
	var err error
	_, err = Calculate(2, 0, "/")
	if err == nil {
		t.Error("Error is nil")
	}
}
