package satsolver

import (
	"fmt"
	"sort"
)

// Solve returns an Interpretation under which the formula is satisfiable
// or an error if the formula is unsatisfiable
func Solve(f Formula, heuristic func(Formula) *Variable) (Interpretation, error) {
	// reset all variables
	for _, variable := range f.V {
		variable.Assigned = false
	}

	nextF, err := reduce(f)
	if err != nil {
		return nil, err
	}

	// the formula was solved by reducing it all the way down
	if len(nextF.C) == 0 {
		return FromVariablesMap(f.V), nil
	}

	// start solving recursively
	err = solveStep(nextF, heuristic)
	if err != nil {
		return nil, err
	}

	return FromVariablesMap(f.V), nil
}

// solveStep returns an error if unsatisfied
func solveStep(f Formula, heuristic func(Formula) *Variable) error {
	variable := heuristic(f)
	if variable == nil {
		return fmt.Errorf("unsatisfiable")
	}

	variable.Assign(true)
	positiveF, err := bcp(f, Literal{
		Negated: false,
		V:       variable,
	})
	// check if we solved or need to recurse
	if err == nil {
		if len(positiveF.C) == 0 {
			return nil
		}

		// recurse if there is no error in this branch it's solved
		err = solveStep(positiveF, heuristic)
		if err == nil {
			return nil
		}
	}

	variable.Assign(false)
	negativeF, err := bcp(f, Literal{
		Negated: true,
		V:       variable,
	})
	// check if we solved or need ot recurse
	if err == nil {
		if len(negativeF.C) == 0 {
			return nil
		}

		// recurse if there is no error in this branch it's solved
		err = solveStep(negativeF, heuristic)
		if err == nil {
			return nil
		}
	}


	// neither branch resulted in a solution reset the variable used
	variable.Assigned = false
	return fmt.Errorf("unsatisfiable")
}

// bcp applies rules to reduce the formula down
// unit rule a ^ (a v b) = T
// tautology (a v ~a) = T
// contradiction a ^ ~a = F
// if there is a contradiction an error is returned
// TODO: remove duplicates with in a clause
func reduce(f Formula) (Formula, error) {
	nextF := CloneFormula(f)
	toRemove := []int{}
	for i, clause := range f.C {
		// unit clause
		if len(clause.L) == 1 {
			unit := clause.L[0]
			unit.V.Assign(!unit.Negated)
			toRemove = append(toRemove, i)

			for j, c := range f.C {
				if i != j {
					if len(c.L) == 1 {
						if c.L[0].V == unit.V {
							// check for contradiction
							if c.L[0].Negated != unit.Negated {
								return Formula{}, fmt.Errorf("contradiction on the variable %s", unit.V.Name)
							} else { // remove duplicate
								toRemove = append(toRemove, j)
							}
						}
					} else {
						// check for the unit rule
						removedLiteralCount := 0
						for k, literal := range c.L {
							if literal.V == unit.V && literal.Negated == unit.Negated {
								toRemove = append(toRemove, j)
								break
							}

							// remove the opposite of literal from clause
							if literal.V == unit.V && literal.Negated != unit.Negated {
								nextF.C[j].L = append(nextF.C[j].L[:k-removedLiteralCount], nextF.C[j].L[k+1-removedLiteralCount:]...)
								removedLiteralCount++

								if len(nextF.C[j].L) == 0 {
									// if there are no literals but the clause wasn't removed we failed
									return Formula{}, fmt.Errorf("unsatisfiable")
								}
							}
						}
					}
				}
			}
		}

		// tautology
		for _, literal := range clause.L {
			found := false
			for _, l := range clause.L {
				if literal.V == l.V && literal.Negated != l.Negated {
					toRemove = append(toRemove, i)
					found = true
					break
				}
			}
			if found {
				break
			}
		}
	}

	// remove clauses from formula
	sort.Ints(toRemove)
	for i, cIndex := range toRemove {
		nextF.C = append(nextF.C[:cIndex-i], nextF.C[cIndex+1-i:]...)
	}

	return nextF, nil
}

// bcp reduces the given formula using the unit rule
func bcp(f Formula, unit Literal) (Formula, error) {
	nextF := CloneFormula(f)
	removedClauseCount := 0
	for i, clause := range f.C {
		removedLiteralCount := 0
		for j, literal := range clause.L {
			// remove clauses the contain the literal
			if literal.V == unit.V && literal.Negated == unit.Negated {
				nextF.C = append(nextF.C[:i-removedClauseCount], nextF.C[i+1-removedClauseCount:]...)
				removedClauseCount++
				break
			}

			// remove the opposite literal
			if literal.V == unit.V && literal.Negated != unit.Negated {
				nextF.C[i-removedClauseCount].L = append(nextF.C[i-removedClauseCount].L[:j-removedLiteralCount], nextF.C[i-removedClauseCount].L[j+1-removedLiteralCount:]...)
				removedLiteralCount++

				if len(nextF.C[i-removedClauseCount].L) == 0 {
					// if there are no literals but the clause wasn't removed we failed
					return Formula{}, fmt.Errorf("unsatisfiable")
				}
			}
		}
	}

	return nextF, nil
}