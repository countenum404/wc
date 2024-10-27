package counter

import (
	"os"
	"testing"
)

const (
	path  = "testFiles/test.txt"
	path1 = "testFiles/test1.txt"
	path2 = "testFiles/test2.txt"
)

func TestFile(t *testing.T) {
	charsToBe, linesToBe, wordsToBe := 4, 2, 3

	file, err := os.Open(path)
	if err != nil {
		t.Error(err)
	}

	chars, lines, words := NewTextFileCounter(file).Count()
	if chars != charsToBe {
		t.Error("Chars aren't equals", chars, charsToBe)
	}

	if lines != linesToBe {
		t.Error("Lines aren't equals", lines, linesToBe)
	}

	if words != wordsToBe {
		t.Error("Words aren't equals", words, wordsToBe)
	}
}

func TestFile1(t *testing.T) {
	charsToBe, linesToBe, wordsToBe := 5, 2, 3

	file, err := os.Open(path1)
	if err != nil {
		t.Error(err)
	}

	chars, lines, words := NewTextFileCounter(file).Count()
	if chars != charsToBe {
		t.Error("Chars aren't equals", chars, charsToBe)
	}

	if lines != linesToBe {
		t.Error("Lines aren't equals", lines, linesToBe)
	}

	if words != wordsToBe {
		t.Error("Words aren't equals", words, wordsToBe)
	}
}

func TestFile2(t *testing.T) {
	charsToBe, linesToBe, wordsToBe := 6, 2, 3

	file, err := os.Open(path2)
	if err != nil {
		t.Error(err)
	}

	chars, lines, words := NewTextFileCounter(file).Count()
	if chars != charsToBe {
		t.Error("Chars aren't equals", chars, charsToBe)
	}

	if lines != linesToBe {
		t.Error("Lines aren't equals", lines, linesToBe)
	}

	if words != wordsToBe {
		t.Error("Words aren't equals", words, wordsToBe)
	}
}
