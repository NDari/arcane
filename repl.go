package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	l := NewLexer()
	rd := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("incantation ::= ")
		str, _ := rd.ReadString('\n')
		if strings.TrimSpace(str) == ":q" {
			fmt.Println("Goodbye!")
			break
		}
		l.SetInput(str)
		for lex := l.NextLexeme(); lex.Type != EOF; lex = l.NextLexeme() {
			fmt.Println(lex)
		}
	}
}
