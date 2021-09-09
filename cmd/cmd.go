package cmd

import (
	"bufio"
	"fmt"
	"git.kanersps.pw/loop/evaluator"
	"git.kanersps.pw/loop/lexer"
	"git.kanersps.pw/loop/models/object"
	"git.kanersps.pw/loop/parser"
	"os"
)

func Execute() {
	// REPL here for testing while developing
	reader := bufio.NewScanner(os.Stdin)

	env := object.CreateEnvironment()

	// TODO: Support up arrow to get last command
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

		if len(p.Errors) > 0 {
			for _, err := range p.Errors {
				fmt.Println(err)
			}
		} else {
			evaluated := evaluator.Eval(program, env)

			fmt.Println(evaluated.Inspect())
		}
	}
}
