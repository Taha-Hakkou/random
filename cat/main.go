package main

import (
	"os"
	"io"
	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		for _, filename := range args {
			data, err := os.ReadFile(filename)
			if err != nil {
				printStr("ERROR: " + err.Error())
				os.Exit(1) //
			}
			printStr(string(data))
		}
	} else {
		for os.Stdin.Read('\n') {
			bytes, _ := io.ReadFull(os.Stdin)
			printStr(string(bytes))
			z01.PrintRune('\n')
		}
	}
}
