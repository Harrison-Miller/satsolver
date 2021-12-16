package satsolver

import (
	"fmt"
	"github.com/alecthomas/participle"
	"io"
	"strings"
)

// MustParse panics if Parse returns an error
func MustParse(s string) Formula {
	r := strings.NewReader(s)
	f, err := Parse(r)
	if err != nil {
		panic(err)
	}

	return f
}

// Parse parses a CNF formula and returns an AST representing it
// example: a ^ (~b v c) ^ ~d
func Parse(r io.Reader) (Formula, error) {
	parser, err := participle.Build(&Formula{})
	if err != nil {
		return Formula{}, fmt.Errorf("error building grammer: %w", err)
	}

	var f Formula
	err = parser.Parse(r, &f)
	if err != nil {
		return Formula{}, fmt.Errorf("error parsing input: %w", err)
	}
	
	CleanFormula(&f)

	return f, nil
}

func CleanFormula(f *Formula) {
	varsMap := map[string]*Variable{}
	for ci, clause := range f.C {
		for li, literal := range clause.L {
			if v, ok := varsMap[literal.V.Name]; ok {
				// fix the literal in clause
				f.C[ci].L[li].V = v
			} else {
				varsMap[literal.V.Name] = literal.V
			}
		}
	}
	f.V = varsMap
}