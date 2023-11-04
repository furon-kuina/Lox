package main

type Token struct {
	TokenType
	lexeme     string
	literal    any
	line       uint
	begin, end uint
}

func NewToken(tt TokenType, lexeme string, line, begin, end uint) Token {
	return Token{
		TokenType: tt,
		lexeme:    lexeme,
		line:      line,
		begin:     begin,
		end:       end,
	}
}

type TokenType uint

const (
	_ TokenType = iota

	// Literals
	Identifier
	String
	Number

	// Keywords
	And
	Class
	Else
	False
	Fun
	For
	If
	Nil
	Or
	Print
	Return
	Super
	This
	True
	Var
	While

	// Single Character
	LeftParen
	RightParen
	LeftBrace
	RightBrace
	Comma
	Dot
	Minus
	Plus
	Semicolon
	Slash
	Star

	Bang
	BangEqual
	Equal
	EqualEqual
	Greater
	GreaterEqual
	Less
	LessEqual

	Eof
)
