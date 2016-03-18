package main

import (
	"bufio"
	"fmt"
	"io"
)

type Token int

const (
	ILLEGAL Token = iota
	EOF
	WS
	IDENT
	ASTERISK
	COMMA
	SELECT
	FROM
)

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

var eof = rune(0)

type Scanner struct {
	r *bufio.Reader
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{bufio.NewReader(r)}
}

func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

func (s *Scanner) unread() { _ = s.r.UnreadRune() }

func (s *Scanner) Scan() (tok Token, lit string) {
	ch := s.read()

	if isLetter(ch) {
		s.unread()
		return s.scanWhiteSpace()
	} else {
		s.unread()
		return s.scanIdent()
	}
	switch ch {
	case eof:
		return EOF, ""
	case '*':
		return ASTERISK, string(ch)
	case ',':
		return COMMA, string(ch)
	}

	return ILLEGAL, string(ch)
}

func (s *Scanner) scanWhiteSpace() (tok Token, lit string) {
}
