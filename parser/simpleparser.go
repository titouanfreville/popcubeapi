// CAUTION: Generated file - DO NOT EDIT.

package parser

import __yyfmt__ "fmt"

import (
	"bufio"
	"fmt"
	"os"
)

type yySymType struct {
	yys   int
	value string
}

type yyXError struct {
	state, xsym int
}

const (
	yyDefault = 57350
	yyEofCode = 57344
	CODE      = 57347
	PING      = 57349
	QUOTE     = 57348
	URL       = 57346
	yyErrCode = 57345

	yyMaxDepth = 200
	yyTabOfs   = -7
)

var (
	yyXLAT = map[int]int{
		10:    0, // '\n' (5x)
		57344: 1, // $end (3x)
		57347: 2, // CODE (2x)
		57349: 3, // PING (2x)
		57348: 4, // QUOTE (2x)
		57351: 5, // text (2x)
		57352: 6, // texts (2x)
		57346: 7, // URL (2x)
		57350: 8, // $default (0x)
		57345: 9, // error (0x)
	}

	yySymNames = []string{
		"'\\n'",
		"$end",
		"CODE",
		"PING",
		"QUOTE",
		"text",
		"texts",
		"URL",
		"$default",
		"error",
	}

	yyTokenLiteralStrings = map[int]string{}

	yyReductions = map[int]struct{ xsym, components int }{
		0: {0, 1},
		1: {6, 3},
		2: {6, 2},
		3: {5, 1},
		4: {5, 1},
		5: {5, 1},
		6: {5, 1},
	}

	yyXErrors = map[yyXError]string{}

	yyParseTab = [9][]uint8{
		// 0
		{2: 11, 13, 12, 9, 8, 10},
		{1: 7},
		{14},
		{4},
		{3},
		// 5
		{2},
		{1},
		{1: 5, 11, 13, 12, 9, 15, 10},
		{1: 6},
	}
)

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyLexerEx interface {
	yyLexer
	Reduced(rule, state int, lval *yySymType) bool
}

func yySymName(c int) (s string) {
	x, ok := yyXLAT[c]
	if ok {
		return yySymNames[x]
	}

	if c < 0x7f {
		return __yyfmt__.Sprintf("'%c'", c)
	}

	return __yyfmt__.Sprintf("%d", c)
}

func yylex1(yylex yyLexer, lval *yySymType) (n int) {
	n = yylex.Lex(lval)
	if n <= 0 {
		n = yyEofCode
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("\nlex %s(%#x %d), lval: %+v\n", yySymName(n), n, n, lval)
	}
	return n
}

func yyParse(yylex yyLexer) int {
	const yyError = 9

	yyEx, _ := yylex.(yyLexerEx)
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, 200)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yyerrok := func() {
		if yyDebug >= 2 {
			__yyfmt__.Printf("yyerrok()\n")
		}
		Errflag = 0
	}
	_ = yyerrok
	yystate := 0
	yychar := -1
	var yyxchar int
	var yyshift int
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
		var ok bool
		if yyxchar, ok = yyXLAT[yychar]; !ok {
			yyxchar = len(yySymNames) // > tab width
		}
	}
	if yyDebug >= 4 {
		var a []int
		for _, v := range yyS[:yyp+1] {
			a = append(a, v.yys)
		}
		__yyfmt__.Printf("state stack %v\n", a)
	}
	row := yyParseTab[yystate]
	yyn = 0
	if yyxchar < len(row) {
		if yyn = int(row[yyxchar]); yyn != 0 {
			yyn += yyTabOfs
		}
	}
	switch {
	case yyn > 0: // shift
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		yyshift = yyn
		if yyDebug >= 2 {
			__yyfmt__.Printf("shift, and goto state %d\n", yystate)
		}
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	case yyn < 0: // reduce
	case yystate == 1: // accept
		if yyDebug >= 2 {
			__yyfmt__.Println("accept")
		}
		goto ret0
	}

	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			if yyDebug >= 1 {
				__yyfmt__.Printf("no action for %s in state %d\n", yySymName(yychar), yystate)
			}
			msg, ok := yyXErrors[yyXError{yystate, yyxchar}]
			if !ok {
				msg, ok = yyXErrors[yyXError{yystate, -1}]
			}
			if !ok && yyshift != 0 {
				msg, ok = yyXErrors[yyXError{yyshift, yyxchar}]
			}
			if !ok {
				msg, ok = yyXErrors[yyXError{yyshift, -1}]
			}
			if yychar > 0 {
				ls := yyTokenLiteralStrings[yychar]
				if ls == "" {
					ls = yySymName(yychar)
				}
				if ls != "" {
					switch {
					case msg == "":
						msg = __yyfmt__.Sprintf("unexpected %s", ls)
					default:
						msg = __yyfmt__.Sprintf("unexpected %s, %s", ls, msg)
					}
				}
			}
			if msg == "" {
				msg = "syntax error"
			}
			yylex.Error(msg)
			Nerrs++
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				row := yyParseTab[yyS[yyp].yys]
				if yyError < len(row) {
					yyn = int(row[yyError]) + yyTabOfs
					if yyn > 0 { // hit
						if yyDebug >= 2 {
							__yyfmt__.Printf("error recovery found error shift in state %d\n", yyS[yyp].yys)
						}
						yystate = yyn /* simulate a shift of "error" */
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery failed\n")
			}
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yySymName(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}

			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	r := -yyn
	x0 := yyReductions[r]
	x, n := x0.xsym, x0.components
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= n
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	exState := yystate
	yystate = int(yyParseTab[yyS[yyp].yys][x]) + yyTabOfs
	/* reduction by production r */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce using rule %v (%s), and goto state %d\n", r, yySymNames[x], yystate)
	}

	switch r {
	case 1:
		{
			fmt.Printf(" \t%T\n", yyS[yypt-2].value)
		}
	case 2:
		{
			fmt.Printf(" \t%T\n", yyS[yypt-1].value)
		}
	case 3:
		{
			yyVAL.value = addUrl(yyS[yypt-0].value, yyVAL.value)
		}
	case 4:
		{
			yyVAL.value = addCode(yyS[yypt-0].value, yyVAL.value)
		}
	case 5:
		{
			yyVAL.value = addQuote(yyS[yypt-0].value, yyVAL.value)
		}
	case 6:
		{
			yyVAL.value = addPing(yyS[yypt-0].value, yyVAL.value)
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}

func addUrl(url string, fullText string) string {
	url = "<a href=\"" + url + "\" target=\"_blank\">" + url + "</a>"
	return url + fullText
}

func addCode(code string, fullText string) string {
	code = "<span class=\"code\">" + code + "</span>"
	var res = code + fullText
	return res
}

func addQuote(quote string, fullText string) string {
	quote = "<quote>" + quote + "</quote>"
	return quote + fullText
}

func addPing(ping string, fullText string) string {
	ping = "<span class=\"ping\">" + ping + "</span>"
	return ping + fullText
}

// type (
//   text interface{}
//   URL string
//   CODE string
//   QUOTE string
//   PING string
// )

func DoParse() {
	os.Exit(yyParse(newLexer(bufio.NewReader(os.Stdin))))
}
