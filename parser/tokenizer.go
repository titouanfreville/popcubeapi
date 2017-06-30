// CAUTION: Generated file - DO NOT EDIT.

package parser

import (
	"bufio"
	"log"
)

type yylexer struct {
	src     *bufio.Reader
	buf     []byte
	empty   bool
	current byte
}

func newLexer(src *bufio.Reader) (y *yylexer) {
	y = &yylexer{src: src}
	if b, err := src.ReadByte(); err == nil {
		y.current = b
	}
	return
}

func (y *yylexer) getc() byte {
	if y.current != 0 {
		y.buf = append(y.buf, y.current)
	}
	y.current = 0
	if b, err := y.src.ReadByte(); err == nil {
		y.current = b
	}
	return y.current
}

func (y yylexer) Error(e string) {
	log.Fatal(e)
}

func (y *yylexer) Lex(lval *yySymType) int {
	// var err error
	c := y.current
	if y.empty {
		c, y.empty = y.getc(), false
	}

yystate0:

	y.buf = y.buf[:0]

	goto yystart1

	goto yystate0 // silence unused label error
	goto yystate1 // silence unused label error
yystate1:
	c = y.getc()
yystart1:
	switch {
	default:
		goto yyrule3
	case c == '#' || c == '%' || c == '+' || c == '-' || c >= '0' && c <= ':' || c == '=' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'g' || c >= 'i' && c <= 'z' || c == '~':
		goto yystate2
	case c == '.':
		goto yystate3
	case c == '>':
		goto yystate4
	case c == '@':
		goto yystate6
	case c == '`':
		goto yystate7
	case c == 'h':
		goto yystate9
	}

yystate2:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c == '#' || c == '%' || c == '+' || c == '-' || c >= '0' && c <= ':' || c == '=' || c >= '@' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '~':
		goto yystate2
	case c == '.':
		goto yystate3
	}

yystate3:
	c = y.getc()
	switch {
	default:
		goto yyrule1
	case c == '#' || c == '%' || c == '+' || c >= '-' && c <= ':' || c == '=' || c >= '@' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '~':
		goto yystate3
	}

yystate4:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c == '\n':
		goto yystate5
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= 'ÿ':
		goto yystate4
	}

yystate5:
	c = y.getc()
	switch {
	default:
		goto yyrule3
	case c == '>':
		goto yystate4
	}

yystate6:
	c = y.getc()
	switch {
	default:
		goto yyrule4
	case c == '#' || c == '%' || c == '+' || c == '-' || c >= '0' && c <= ':' || c == '=' || c == '@' || c == '_' || c == '~':
		goto yystate2
	case c == '.':
		goto yystate3
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate6
	}

yystate7:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c == '`':
		goto yystate8
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '_' || c >= 'a' && c <= 'ÿ':
		goto yystate7
	}

yystate8:
	c = y.getc()
	switch {
	default:
		goto yyrule2
	case c == '`':
		goto yystate8
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '_' || c >= 'a' && c <= 'ÿ':
		goto yystate7
	}

yystate9:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c == '#' || c == '%' || c == '+' || c == '-' || c >= '0' && c <= ':' || c == '=' || c >= '@' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '~':
		goto yystate2
	case c == '.':
		goto yystate3
	case c == 't':
		goto yystate10
	}

yystate10:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c == '#' || c == '%' || c == '+' || c == '-' || c >= '0' && c <= ':' || c == '=' || c >= '@' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '~':
		goto yystate2
	case c == '.':
		goto yystate3
	case c == 't':
		goto yystate11
	}

yystate11:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c == '#' || c == '%' || c == '+' || c == '-' || c >= '0' && c <= ':' || c == '=' || c >= '@' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'o' || c >= 'q' && c <= 'z' || c == '~':
		goto yystate2
	case c == '.':
		goto yystate3
	case c == 'p':
		goto yystate12
	}

yystate12:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c == '#' || c == '%' || c == '+' || c == '-' || c >= '0' && c <= '9' || c == '=' || c >= '@' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z' || c == '~':
		goto yystate2
	case c == '.':
		goto yystate3
	case c == ':':
		goto yystate13
	case c == 's':
		goto yystate15
	}

yystate13:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c == '#' || c == '%' || c == '+' || c == '-' || c >= '0' && c <= ':' || c == '=' || c >= '@' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '~':
		goto yystate2
	case c == '.':
		goto yystate3
	case c == '/':
		goto yystate14
	}

yystate14:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c == '/':
		goto yystate2
	}

yystate15:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c == '#' || c == '%' || c == '+' || c == '-' || c >= '0' && c <= '9' || c == '=' || c >= '@' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '~':
		goto yystate2
	case c == '.':
		goto yystate3
	case c == ':':
		goto yystate13
	}

yyrule1: // {url}
	{
		return URL
	}
yyrule2: // `.*`|```.*```
	{
		return CODE
	}
yyrule3: // (\>.*\n)*
	{
		return QUOTE
	}
yyrule4: // @[a-zA-Z]*
	{
		return PING
	}
	panic("unreachable")

	goto yyabort // silence unused label error

yyabort: // no lexem recognized
	y.empty = true
	return int(c)
}
