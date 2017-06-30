%{
package parser

import ( 
  "bufio"
  "fmt"
  "os"
)
%}

%union{
  value string
}

%token URL CODE QUOTE PING

%type <value> URL, CODE, QUOTE, PING, text
%%
texts: text '\n' texts {fmt.Printf(" \t%T\n", $1)} | text '\n' {fmt.Printf(" \t%T\n", $1)};
text:
  URL {$$=addUrl($1, $$)} |
  CODE {$$=addCode($1, $$)} |
  QUOTE {$$=addQuote($1, $$)} |
  PING {$$=addPing($1, $$)}
  ;

%%
func addUrl(url string, fullText string) string {
  url= "<a href=\""+ url +"\" target=\"_blank\">"+url+"</a>" 
  return url + fullText
}

func addCode(code string, fullText string) string {
  code= "<span class=\"code\">"+code+"</span>" 
  var res = code + fullText
  return res
}

func addQuote(quote string, fullText string) string {
  quote= "<quote>"+quote+"</quote>" 
  return quote + fullText
}

func addPing(ping string, fullText string) string {
  ping= "<span class=\"ping\">"+ping+"</span>" 
  return ping +fullText
}

// type (
//   text interface{}
//   URL string
//   CODE string
//   QUOTE string
//   PING string
// )

func  DoParse() {
  os.Exit(yyParse(newLexer(bufio.NewReader(os.Stdin))));
}