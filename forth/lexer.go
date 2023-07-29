package forth

type Lexer interface {
	// NextToken returns a next token.
	NextToken() Token
	Reset()
}

type lexer struct {
	input        string
	curPos       int
	readPosition int
	char         byte
}

// Resets the lexer state back to its original one
func (l *lexer) Reset() {
	l.curPos = 0
	l.readPosition = 0
	l.char = 0
	l.readChar()
	return
}

// New returns a new Lexer
func NewLex(input string) Lexer {
	l := &lexer{input: input}
	l.readChar()
	return l
}

func (l *lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readPosition]
	}
	l.curPos = l.readPosition
	l.readPosition++
}

func (l *lexer) NextToken() Token {
	l.skipWhitespace()

	// skip comments
	if l.char == '/' && l.peekChar() == '/' {
		l.skipComment()
	}

	var tok Token
	switch l.char {
	case '.':
		if l.peekChar() == '"' {
			l.readChar()
			l.readChar()
			fullString := ".\""
			for l.char != '"' {
				fullString += string(l.char)
				l.readChar()
			}
			fullString += "\""
			tok.Type = DOTPRINT
			tok.Literal = fullString
		} else {
			tok = newToken(DOT, l.char)
		}
	case '=':
		tok = newToken(EQ, l.char)
	case ':':
		tok = newToken(COLON, l.char)
	case ';':
		tok = newToken(SEMICOLON, l.char)
	case '+':
		tok = newToken(PLUS, l.char)
	case '-':
		tok = newToken(MINUS, l.char)
	case '*':
		tok = newToken(ASTERISK, l.char)
	case '<':
		tok = newToken(LT, l.char)
	case '>':
		tok = newToken(GT, l.char)
	case 0:
		tok.Literal = ""
		tok.Type = EOF
	default:
		if isDigit(l.char) {
			return l.readNumberToken()
		}

		if isLetter(l.char) {
			tok.Literal = l.readIdent()
			tok.Type = LookupIdent(tok.Literal)
			return tok
		}

		tok = newToken(ILLEGAL, l.char)
	}

	l.readChar()
	return tok
}

func (l *lexer) skipWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}

func (l *lexer) skipComment() {
	for l.char != '\n' && l.char != '\r' {
		l.readChar()
	}
	l.skipWhitespace()
}

func (l *lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *lexer) readString() string {
	position := l.curPos + 1
	for {
		l.readChar()
		if l.char == '"' || l.char == 0 {
			break
		}
	}
	return l.input[position:l.curPos]
}

func (l *lexer) read(checkFn func(byte) bool) string {
	position := l.curPos
	for checkFn(l.char) {
		l.readChar()
	}
	return l.input[position:l.curPos]
}

func (l *lexer) readIdent() string {
	return l.read(isLetter)
}

func (l *lexer) readNumber() string {
	return l.read(isDigit)
}

// TODO: Return float and turn ints into floats
func (l *lexer) readNumberToken() Token {
	intPart := l.readNumber()
	return Token{
		Type:    NUMBER,
		Literal: intPart,
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func newToken(tokenType TokenType, ch byte) Token {
	return Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}
