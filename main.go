package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	Start(os.Stdin, os.Stdout)
}

func Start(in io.Reader, out io.Writer) {
	const PROMPT = ">> "
	scanner := bufio.NewScanner(in)
	for {
		fmt.Println("Escribe x si quieres terminar el programa")
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		input := scanner.Text()
		if input == "X" || input == "x" {
			fmt.Println("Se acabo")
			break
		}
	}
}
