package processor

import "testing"

func TestHexConversion(t *testing.T) {
	input := "Simply add 42 (hex)."
	expected := "Simply add 66."

	result := Process(input)

	if result != expected {
		t.Errorf("Expected: %q\nGot:      %q", expected, result)
	}
}

func TestBinConversion(t *testing.T) {
	input := "10 (bin)"
	expected := "2"

	result := Process(input)

	if result != expected {
		t.Errorf("Expected: %q\nGot:      %q", expected, result)
	}
}

func TestUpCommand(t *testing.T) {
	input := "hello world (up)"
	expected := "hello WORLD"

	result := Process(input)

	if result != expected {
		t.Errorf("Expected: %q\nGot:      %q", expected, result)
	}
}

func TestCapMultiple(t *testing.T) {
	input := "it was the age of foolishness (cap, 6)"
	expected := "It Was The Age Of Foolishness"

	result := Process(input)

	if result != expected {
		t.Errorf("Expected: %q\nGot:      %q", expected, result)
	}
}

func TestLowMultiple(t *testing.T) {
	input := "IT WAS THE (low, 3) winter"
	expected := "it was the winter"

	result := Process(input)

	if result != expected {
		t.Errorf("Expected: %q\nGot:      %q", expected, result)
	}
}

func TestArticleCorrection(t *testing.T) {
	input := "There is a apple."
	expected := "There is an apple."

	result := Process(input)

	if result != expected {
		t.Errorf("Expected: %q\nGot:      %q", expected, result)
	}
}

func TestPunctuationSpacing(t *testing.T) {
	input := "Hello , world !"
	expected := "Hello, world!"

	result := Process(input)

	if result != expected {
		t.Errorf("Expected: %q\nGot:      %q", expected, result)
	}
}

func TestEllipsis(t *testing.T) {
	input := "This is ... amazing !"
	expected := "This is... amazing!"

	result := Process(input)

	if result != expected {
		t.Errorf("Expected: %q\nGot:      %q", expected, result)
	}
}

func TestFullExample(t *testing.T) {
	input := "Simply add 42 (hex) and 10 (bin)."
	expected := "Simply add 66 and 2."

	result := Process(input)

	if result != expected {
		t.Errorf("Expected: %q\nGot:      %q", expected, result)
	}
}
