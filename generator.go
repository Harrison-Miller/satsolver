package satsolver

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateUniform3SAT generates a random 3SAT problem
// m is the number of clauses
// n is the number of variables
func GenerateUniform3SAT(m int, n int, seed *int64) (Formula, error) {
	// seed randomness
	var r *rand.Rand
	if seed != nil {
		r = rand.New(rand.NewSource(*seed))
	} else {
		r = rand.New(rand.NewSource(time.Now().UnixNano()))
	}

	// create variables
	variables := make([]*Variable, n)
	for i := 0; i < n; i++ {
		nameI := i + 1
		var name string
		for nameI > 0 {
			modulo := (nameI - 1) % 26
			name = fmt.Sprintf("%s%s", string(byte(97+modulo)), name)
			nameI = (nameI - modulo) / 26
		}

		variables[i] = &Variable{
			Name:     name,
			Value:    false,
			Assigned: false,
		}
	}

	// create clauses and formula
	var f Formula
	for i := 0; i < m; i++ {
		c := Clause{
			L: []Literal{
				{
					Negated: toBool(r.Intn(2)),
					V:      variables[r.Intn(len(variables))],
				},
				{
					Negated: toBool(r.Intn(2)),
					V: variables[r.Intn(len(variables))],
				},
				{
					Negated: toBool(r.Intn(2)),
					V: variables[r.Intn(len(variables))],
				},
			},
		}
		f.C = append(f.C, c)
	}

	CleanFormula(&f)

	return f, nil
}

func toBool(x int) bool {
	return x > 0
}