package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/NDari/arcane/reader"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	counter := 1
	for {
		lines := ""
		for scanner.Scan() {
			line := scanner.Text()
			if lines == "" {
				line = strings.TrimLeft(line, " ")
				if !strings.HasPrefix(line, "(") {
					fmt.Println("expression must start with '('")
					break
				}
				lines += line
			} else {
				lines = lines + "\n" + line
			}
			if reader.ValidExprString(lines) {
				break
			}
		}

		r := reader.NewReader(lines)
		anys, err := r.ReadAll()
		if err != nil {
			fmt.Println(err, "\n")
			continue
		}
		fmt.Printf("$%d = ", counter)
		iter := anys.ToIterable()
		for {
			if iter.HasNext() {
				item := iter.Next()
				fmt.Print(item.Repr() + " ")
			} else {
				break
			}
		}
		fmt.Println()
		counter++
	}
}
