package forth

type TokenType string

const (
	//Special Types
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	IDENT   = "IDENT"
	NUMBER  = "NUMBER"
	// Arithmetic Operators
	PLUS     = "PLUS"     // +
	MINUS    = "MINUS"    // -
	ASTERISK = "ASTERISK" // *
	MOD      = "MOD"      // mod
	// Stack Operations
	DUP  = "DUP"  // dup
	DROP = "DROP" // drop
	SWAP = "SWAP" // swap
	ROT  = "ROT"  // rot
	OVER = "OVER" // over
	// Logical Operators
	AND    = "AND"
	OR     = "OR"
	INVERT = "INVERT"
	GT     = "GT"   // >
	LT     = "LT"   // <
	EQ     = "EQ"   // =
	IF     = "IF"   // if
	ELSE   = "ELSE" // else
	THEN   = "THEN" // then
	// Function Definition
	COLON     = "COLON"     // :
	SEMICOLON = "SEMICOLON" // ;
	// Loops
	DO    = "DO"    // do
	LOOP  = "LOOP"  // loop
	INDEX = "INDEX" // idx
	// Output
	DOTPRINT = "DOTPRINT" // ."
	EMIT     = "EMIT"     // emit
	CR       = "CR"       // cr
	DOT      = "DOT"      // .
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"do":     DO,
	"loop":   LOOP,
	"idx":    INDEX,
	"if":     IF,
	"else":   ELSE,
	"then":   THEN,
	"mod":    MOD,
	"cr":     CR,
	"emit":   EMIT,
	"rot":    ROT,
	"over":   OVER,
	"swap":   SWAP,
	"drop":   DROP,
	"dup":    DUP,
	"and":    AND,
	"invert": INVERT,
	"or":     OR,
}

// Lookup ident checks to see if the given identifier is
// a keyword or not. If it is then return it's type otherwise
// return IDENT type.
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

func CombineTokens(tokList []Token) string {
	combined := ""
	for _, tok := range tokList {
		combined += tok.Literal + " "
	}
	return combined
}
