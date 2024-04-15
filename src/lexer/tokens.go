package lexer

import "fmt"

type TokenKind int

const (
	EOF TokenKind = iota
	NUMBER
	STRING
	IDENTIFIER

	OPEN_BRACKET
	CLOSE_BRACKET
	OPEN_CURLY
	CLOSE_CURLY
	OPEN_PAREN
	CLOSE_PAREN

	ASSIGNMENT
	EQUALS
	NOT
	NOT_EQUALS

	LESS
	LESS_EQUALS
	GREATER
	GREATER_EQUALS

	OR
	AND

	DOT
	DOT_DOT
	SEMI_COLON
	COLON
	QUESTION
	COMMA

	PLUS_PLUS
	MINUS_MINUS
	PLUS_EQUALS
	MINUS_EQUALS

	PLUS
	DASH
	SLASH
	STAR
	PERCENT

	// Reserved Keywords
	LET
	CONST
	CLASS
	NEW
	IMPORT
	FROM
	FN
	IF
	ELIF
	ELSE
	FOREACH
	WHILE
	FOR
	EXPORT
	TYPEOF
	IN
)

var reserved_lu map[string]TokenKind = map[string]TokenKind{
	"let":     LET,
	"const":   CONST,
	"class":   CLASS,
	"new":     NEW,
	"import":  IMPORT,
	"from":    FROM,
	"fn":      FN,
	"if":      IF,
	"elif":    ELIF,
	"else":    ELSE,
	"foreach": FOREACH,
	"while":   WHILE,
	"for":     FOR,
	"export":  EXPORT,
	"typeof":  TYPEOF,
	"in":      IN,
}

type Token struct {
	Kind  TokenKind
	Value string
}

func (token Token) isOneOfMany(expectedTokens ...TokenKind) bool {
	for _, expected := range expectedTokens {
		if expected == token.Kind {
			return true
		}
	}

	return false
}

func (token Token) Debug() {
	if token.isOneOfMany(IDENTIFIER, NUMBER, STRING) {
		fmt.Printf("%s {%s}\n", token.Kind.toString(), token.Value)
	} else {
		fmt.Printf("%s {}\n", token.Kind.toString())
	}
}

func NewToken(kind TokenKind, value string) Token {
	return Token{kind, value}
}

func (kind TokenKind) toString() string {
	names := [...]string{
		"eof",
		"number",
		"string",
		"identifier",
		"open_bracket",
		"close_bracket",
		"open_curly",
		"close_curly",
		"open_paren",
		"close_paren",
		"assignment",
		"equals",
		"not",
		"not_equals",
		"less",
		"less_equals",
		"greater",
		"greater_equals",
		"or",
		"and",
		"dot",
		"dot_dot",
		"semi_colon",
		"colon",
		"question",
		"comma",
		"plus_plus",
		"minus_minus",
		"plus_equals",
		"minus_equals",
		"plus",
		"dash",
		"slash",
		"star",
		"percent",
		"let",
		"const",
		"class",
		"new",
		"import",
		"from",
		"fn",
		"if",
		"elif",
		"else",
		"foreach",
		"while",
		"for",
		"export",
		"typeof",
		"in",
	}
	if kind < EOF || kind > IN {
		return "unknown"
	}
	return names[kind]
}
