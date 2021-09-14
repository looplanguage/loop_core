package cmd

import (
	"bufio"
	"fmt"
	"git.kanersps.pw/loop/evaluator"
	"git.kanersps.pw/loop/lexer"
	"git.kanersps.pw/loop/models/object"
	"git.kanersps.pw/loop/parser"
	"io/ioutil"
	"log"
	"os"
)

func Execute() {
	// Check if file is specified, run if it is
	if len(os.Args) == 2 {
		input, err := ioutil.ReadFile(os.Args[1])

		if err != nil {
			log.Fatal(err)
		}

		env := object.CreateEnvironment("")

		l := lexer.Create(string(input))
		p := parser.Create(l)

		program := p.Parse()

		if len(p.Errors) > 0 {
			for _, err := range p.Errors {
				fmt.Println(err)
			}
		} else {
			evaluated := evaluator.Eval(program, env)

			_, none := evaluated.(*object.Null)

			if !none {
				fmt.Println(evaluated.Inspect())
			}
		}

		os.Exit(0)
	}

	// REPL here for testing while developing
	reader := bufio.NewScanner(os.Stdin)

	env := object.CreateEnvironment("")

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

			_, none := evaluated.(*object.Null)

			if !none {
				fmt.Println(evaluated.Inspect())
			}
		}
	}
}
