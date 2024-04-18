/*
REPL = "Read Eval Print Loop"
This is what we often call the console or interactive mode, it's the mode that starts whenever we type python in the cmd,
ie

$ python
>> 5 + 5
10
*/
package repl

import (
	"bufio"
	"fmt"
	"github.com/negarciacamilo/go_interpreter/lexer"
	"github.com/negarciacamilo/go_interpreter/token"
	"io"
)

const PROMT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	// Read from the input source until a new line
	// Take the just read line and pass it to an instance of our lexer
	// Print all the tokens the lexer gives us until we encounter EOF
	for {
		fmt.Printf(PROMT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
