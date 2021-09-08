package cmd

import (
	"bufio"
	"fmt"
	"git.kanersps.pw/loop/evaluator"
	"git.kanersps.pw/loop/lexer"
	"git.kanersps.pw/loop/parser"
	"os"
)

func Execute() {
	// REPL here for testing while developing
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("# ")

		scanned := reader.Scan()

		if !scanned {
			return
		}

		input := reader.Text()

		l := lexer.Create(input)
		p := parser.Create(l)

		program := p.Parse()

		evaluated := evaluator.Eval(program)

		fmt.Println(evaluated.Inspect())
	}
}
