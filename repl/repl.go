package repl

import (
	"bufio"
	"fmt"
	"io"
	"tyr/evaluator"
	"tyr/lexer"
	"tyr/object"
	"tyr/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

const TYR_FACE = `            __,__
________________.___.__________ 
\__    ___/\__  |   |\______   \
  |    |    /   |   | |       _/
  |    |    \____   | |    |   \
  |____|    / ______| |____|_  /
            \/               \/ 
`

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, TYR_FACE)
	io.WriteString(out, "Woops! We ran into some tyr business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
