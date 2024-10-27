package counter

import (
	"io"
	"os"
)

const (
	SPACE   = 32
	NEWLINE = 10
)

type Counnter interface {
	Count() (int, int, int)
}

type TextFileCounter struct {
	file *os.File
}

func NewTextFileCounter(file *os.File) *TextFileCounter {
	return &TextFileCounter{file: file}
}

func (c *TextFileCounter) Count() (int, int, int) {
	chars, lines, words := 0, 0, 0
	b := make([]byte, 1)

	insideWord := false

	for {
		_, err := c.file.Read(b)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		char := b[0]
		if char == SPACE {
			chars++
			if insideWord {
				words++
			}
			insideWord = false
		} else if char == NEWLINE {
			lines++
			if insideWord {
				words++
			}
			insideWord = false
		} else {
			insideWord = true
			chars++
		}
	}
	if chars != 0 {
		lines++
	}
	chars--
	if insideWord {
		words++
	}
	return chars, lines, words
}
