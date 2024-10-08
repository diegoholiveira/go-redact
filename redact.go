package redact

import (
	"fmt"
	"io"
	"strings"
	"unicode"
	"unicode/utf8"
)

func Redact(input io.Reader, output io.Writer, expressions []string, redactedValue string) {
	t := NewTrie()
	for _, expression := range expressions {
		t.Insert(strings.ToLower(expression))
	}

	sb := strings.Builder{}
	for {
		var buf [4]byte // A rune can be represented by up to 4 bytes in UTF-8

		_, err := input.Read(buf[:])
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error reading file:", err)
			break
		}

		if utf8.RuneCount(buf[:]) == 1 {
			_, _ = output.Write(buf[:])

			continue
		}

		for _, c := range buf {
			if unicode.IsLetter(rune(c)) || unicode.IsNumber(rune(c)) {
				_, _ = sb.Write([]byte{c})
			} else {
				w := sb.String()

				sb = strings.Builder{}

				if t.Search(strings.ToLower(w)) {
					w = redactedValue
				}

				_, _ = output.Write([]byte(w))
				_, _ = output.Write([]byte{c})
			}
		}
	}
}
