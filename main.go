package main

import (
	"monkeylang/repl"
	"os"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
