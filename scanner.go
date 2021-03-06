package main

import (
	"bufio"
	"errors"
	"io"
)

type Token int

const (
	EOF      Token = iota
	FORWARD        // >
	BACKWARD       // <
	ADD            // +
	SUB            // -
	OUTPUT         // .
	INPUT          // ,
	JMPL           // [
	JMPR           // ]
)

type Scanner struct {
	reader *bufio.Reader

	ch rune // current character
}

func NewScanner(reader *bufio.Reader) *Scanner {
	return &Scanner{reader: reader}
}

func (s *Scanner) next() {
	var err error
	s.ch, _, err = s.reader.ReadRune()
	if err != nil {
		if errors.Is(err, io.EOF) {
			s.ch = -1
			return
		}

		panic(err)
	}
}

func (s *Scanner) Scan() (tok Token) {
	s.next()

	switch s.ch {
	case -1:
		return EOF
	case '>':
		return FORWARD
	case '<':
		return BACKWARD
	case '+':
		return ADD
	case '-':
		return SUB
	case '.':
		return OUTPUT
	case ',':
		return INPUT
	case '[':
		return JMPL
	case ']':
		return JMPR
	default:
		return s.Scan()
	}
}
