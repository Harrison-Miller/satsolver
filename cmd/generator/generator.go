package main

import (
	"fmt"
	"os"
	"satsolver"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		usage()
	}

	m, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(fmt.Errorf("failed to parse m: %w", err))
		os.Exit(-1)
	}

	n, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(fmt.Errorf("failed to parse n: %w", err))
	}

	f, err := satsolver.GenerateUniform3SAT(m, n, nil)
	if err != nil {
		panic(fmt.Errorf("failed to generate 3SAT instance: %w", err))
		return
	}

	fmt.Println(f)
}

func usage() {
	fmt.Printf("%s <m> <n>\n", os.Args[0])
	fmt.Println("m is the number of clauses")
	fmt.Println("n is the number of variables")
	os.Exit(-1)
}

