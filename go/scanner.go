package main

import (
	"fmt"
)

type Scanner struct {
	src                string
	err                error
	line, col, current uint
	start              uint
	tokens             []Token
	lastCh             byte // most recently read character
}

const null byte = '\000'

var keywords = map[string]TokenType{
	"and":    And,
	"class":  Class,
	"else":   Else,
	"false":  False,
	"for":    For,
	"fun":    Fun,
	"if":     If,
	"nil":    Nil,
	"or":     Or,
	"print":  Print,
	"return": Return,
	"super":  Super,
	"this":   This,
	"true":   True,
	"var":    Var,
	"while":  While,
}

func NewScanner(src string) *Scanner {
	nsrc := string(null)
	nsrc += src
	return &Scanner{
		src:     nsrc,
		err:     nil,
		line:    1,
		current: 1,
		start:   1,
		lastCh:  null,
	}
}

func (s *Scanner) reportError(err error) {
	if s.err == nil {
		s.err = fmt.Errorf("Found error at line %d: %v", s.line, err)
	}
}

// scanTokens reads from src, and split the input into tokens
func (s *Scanner) scanTokens() {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}
	s.start = s.current
	s.addToken(Eof)
}

func (s *Scanner) scanToken() {
	ch := s.advance()
	switch ch {
	case '(':
		s.addToken(LeftParen)
	case ')':
		s.addToken(RightParen)
	case '{':
		s.addToken(LeftBrace)
	case '}':
		s.addToken(RightBrace)
	case ',':
		s.addToken(Comma)
	case '.':
		s.addToken(Dot)
	case '-':
		s.addToken(Minus)
	case '+':
		s.addToken(Plus)
	case ';':
		s.addToken(Semicolon)
	case '*':
		s.addToken(Star)
	case '!':
		if s.match('=') {
			s.addToken(BangEqual)
		} else {
			s.addToken(Bang)
		}
	case '=':
		if s.match('=') {
			s.addToken(EqualEqual)
		} else {
			s.addToken(Equal)
		}
	case '<':
		if s.match('=') {
			s.addToken(LessEqual)
		} else {
			s.addToken(Less)
		}
	case '>':
		if s.match('=') {
			s.addToken(GreaterEqual)
		} else {
			s.addToken(Greater)
		}

	case '/':
		if s.match('/') {
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(Slash)
		}

	case ' ':
	case '\r':
	case '\t':
	case '\n':
		s.line++
	case '"':
		s.string()

	default:
		if isDigit(ch) {
			s.number()
		} else if isAlpha(ch) {
			s.identifier()
		} else {
			s.reportError(fmt.Errorf("Unexpected character %c", ch))
		}
	}
}

func (s *Scanner) addToken(tt TokenType) {
	s.tokens = append(s.tokens, NewToken(tt, s.src[s.start:s.current], s.line, s.start, s.current))
}

func (s *Scanner) addTokenRange(tt TokenType, begin, end uint) {
	s.tokens = append(s.tokens, NewToken(tt, s.src[begin:end], s.line, begin, end))
}

func (s *Scanner) match(ch byte) bool {
	// Thanks to the sentinel, we can assume len(s.src) >= 1
	if s.isAtEnd() {
		return false
	}
	if s.src[s.current] == ch {
		s.current++
		return true
	} else {
		return false
	}
}

func (s *Scanner) peek() byte {
	if s.isAtEnd() {
		return null
	}
	return s.src[s.current]
}

func (s *Scanner) advance() byte {
	res := s.src[s.current]
	s.current++
	return res
}

func (s *Scanner) isAtEnd() bool {
	return s.current == uint(len(s.src))
}

func (s *Scanner) string() {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}
	if s.isAtEnd() {
		s.reportError(fmt.Errorf("Unterminated string"))
		return
	}
	s.advance()
	s.addTokenRange(String, s.start, s.current+1)
}

func (s *Scanner) number() {
	for isDigit(s.peek()) {
		s.advance()
	}
	if s.peek() == '.' && isDigit(s.peekNext()) {
		s.advance()
		for isDigit(s.peek()) {
			s.advance()
		}
	}
	s.addToken(Number)
}

func (s *Scanner) peekNext() byte {
	if s.current+1 >= uint(len(s.src)-1) {
		return null
	}
	return s.src[s.current+1]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isAlpha(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isAlphaNumeric(ch byte) bool {
	return isAlpha(ch) || isDigit(ch)
}

func (s *Scanner) identifier() {
	for isAlphaNumeric(s.peek()) {
		s.advance()
	}
	text := s.src[s.start:s.current]
	tt, ok := keywords[text]
	if !ok {
		tt = Identifier
	}
	s.addToken(tt)
}
