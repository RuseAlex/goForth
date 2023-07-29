package forth

import (
	"fmt"
	"log"
	"strconv"
)

var IdentMap = map[string]string{}

type Eval struct {
	stack     *Stack
	lex       Lexer
	wordCount int
	funMap    *map[string]string
}

func (e *Eval) PrintStack() {
	fmt.Println(e.stack.String())
}

func NewEval(input string, stack *Stack) Eval {
	lex := NewLex(input)
	return Eval{
		stack:  stack,
		lex:    lex,
		funMap: &IdentMap,
	}
}

func (e *Eval) Exec() {
	tok := e.lex.NextToken()
	for tok.Type != EOF {
		if tok.Type == ILLEGAL {
			log.Panicln("\nerror: caught an illegal token ", tok)
		}
		if tok.Type == NUMBER {
			tokIntVal, _ := strconv.Atoi(tok.Literal)
			e.stack.Push(tokIntVal)
		}
		if tok.Type == PLUS {
			if err := e.add(); err != nil {
				log.Printf("\nerror: couldn't add because :%s \n", err)
			}
		}
		if tok.Type == MINUS {
			if err := e.subtract(); err != nil {
				log.Printf("\nerror: couldn't subtract because :%s \n", err)
			}
		}
		if tok.Type == ASTERISK {
			if err := e.multiply(); err != nil {
				log.Printf("\nerror: couldn't multiply because :%s \n", err)
			}
		}
		if tok.Type == MOD {
			if err := e.mod(); err != nil {
				log.Printf("\nerror: couldn't modulo because :%s \n", err)
			}
		}
		if tok.Type == OR {
			if err := e.or(); err != nil {
				log.Printf("\nerror: couldn't OR because :%s \n", err)
			}
		}
		if tok.Type == AND {
			if err := e.and(); err != nil {
				log.Printf("\nerror: couldn't AND because :%s \n", err)
			}
		}
		if tok.Type == INVERT {
			if err := e.invert(); err != nil {
				log.Printf("\nerror: couldn't INVERT because :%s \n", err)
			}
		}
		if tok.Type == DOTPRINT {
			e.dotPrint(tok)
		}
		if tok.Type == SWAP {
			if err := e.swap(); err != nil {
				log.Printf("\nerror: couldn't swap because :%s \n", err)
			}
		}
		if tok.Type == ROT {
			if err := e.rot(); err != nil {
				log.Printf("\nerror: couldn't rot because :%s \n", err)
			}
		}
		if tok.Type == OVER {
			if err := e.over(); err != nil {
				log.Printf("\nerror: couldn't over because :%s \n", err)
			}
		}
		if tok.Type == DROP {
			if err := e.drop(); err != nil {
				log.Printf("\nerror: couldn't drop because :%s \n", err)
			}
		}
		if tok.Type == DUP {
			if err := e.dup(); err != nil {
				log.Printf("\nerror: couldn't dup because :%s \n", err)
			}
		}
		if tok.Type == CR {
			if err := e.cr(); err != nil {
				log.Printf("\nerror: couldn't cr because :%s \n", err)
			}
		}
		if tok.Type == EMIT {
			if err := e.emit(); err != nil {
				log.Printf("\nerror: couldn't emit because :%s \n", err)
			}
		}
		if tok.Type == DOT {
			if err := e.dot(); err != nil {
				log.Printf("\nerror: couldn't dot because :%s \n", err)
			}
		}
		if tok.Type == EQ {
			if err := e.eq(); err != nil {
				log.Printf("\nerror: couldn't eq because :%s \n", err)
			}
		}
		if tok.Type == LT {
			if err := e.lt(); err != nil {
				log.Printf("\nerror: couldn't lt because :%s \n", err)
			}
		}
		if tok.Type == GT {
			if err := e.gt(); err != nil {
				log.Printf("\nerror: couldn't gt because :%s \n", err)
			}
		}
		if tok.Type == IDENT {
			if code, ok := IdentMap[tok.Literal]; ok {
				neval := NewEval(code, e.stack)
				neval.wordCount = e.wordCount
				neval.Exec()
			} else {
				log.Fatalln("unrecognized identifier")
			}
		}
		if tok.Type == COLON {
			wordTok := e.lex.NextToken()
			if wordTok.Type != IDENT {
				log.Fatalf("got %s expected an identifier", wordTok.Literal)
			}
			tmpTok := e.lex.NextToken()
			tokList := []Token{}
			for tmpTok.Type != SEMICOLON {
				tokList = append(tokList, tmpTok)
				tmpTok = e.lex.NextToken()
			}
			IdentMap[wordTok.Literal] = CombineTokens(tokList)
		}
		tok = e.lex.NextToken()
	}
}
