package main

import (
	"fmt"
	"github.com/Harrison-Miller/satsolver"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		usage()
	}

	m, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("failed to parse m: ", err)
		os.Exit(-1)
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("failed to parse n: ", err)
		os.Exit(-1)
	}

	f, err := satsolver.GenerateUniform3SAT(m, n, nil)
	if err != nil {
		fmt.Println("failed to generate a valid 3SAT instance: ", err)
		os.Exit(-1)
	}

	fmt.Println(f)
}

func usage() {
	fmt.Printf("%s m n\n", os.Args[0])
	fmt.Println("m is the number of clauses")
	fmt.Println("n is the number of variables")
	os.Exit(-1)
}
