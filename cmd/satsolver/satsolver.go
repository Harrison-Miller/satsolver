package main

import (
	"bufio"
	"fmt"
	"os"
	"satsolver"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		filepath := os.Args[1]
		runFromFile(filepath)
	} else {
		repl()
	}
}

func runFromFile(filepath string) {
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	formula, err := satsolver.Parse(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	I, err := satsolver.Solve(formula, satsolver.JeroslowWang)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("satisfiable with:")
	fmt.Println(I)
}

func repl() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Boolean Satisfiability Shell")
	fmt.Println("---------------------")

	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')

		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		text = strings.TrimSpace(text)

		if strings.ToLower(text) == "quit" {
			os.Exit(0)
		} else if strings.ToLower(text) == "help" {
			fmt.Println("formulas must be in CNF (conjunctive normal form)")
			fmt.Println("example: a ^ (b v ~c) ^ ~d")
			continue
		} else if text == "" {
			continue
		}

		f, err := satsolver.Parse(strings.NewReader(text))
		if err != nil {
			fmt.Println(err)
		} else {
			I, err := satsolver.Solve(f, satsolver.JeroslowWang)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("satisfiable with:")
			fmt.Println(I)
		}

	}
}
