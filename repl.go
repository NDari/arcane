package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/NDari/arcane/reader"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	counter := 1
	for {
		fmt.Print("arcane:: ")
		str, _ := rd.ReadString('\n')
		if strings.TrimSpace(str) == ":q" {
			fmt.Println("Goodbye!")
			break
		}
		r := reader.NewReader(str)
		anys, err := r.Read()
		if err != nil {
			fmt.Println("reader failed:", err)
			continue
		}
		fmt.Printf("$%d = ", counter)
		for i := range anys {
			fmt.Print(anys[i].Repr() + " ")
		}
		fmt.Println()
		counter++
	}
}
