package satsolver

import "fmt"

// Verify returns true if the CNF formula is satisfied under the given interpretation
func Verify(f Formula, I Interpretation) (bool, error) {
	vars := Vars(f)
	for _, variable := range vars {
		if v, ok := I[variable.Name]; ok {
			variable.Value = v
		} else {
			return false, fmt.Errorf("variable %s not assigned", variable.Name)
		}
	}

	for _, clause := range f.C {
		satisfied := false
		for _, literal := range clause.L {
			// find a literal that evaluates to true
			if (!literal.Negated && literal.V.Value) || (literal.Negated && !literal.V.Value) {
				satisfied = true
				break
			}
		}

		if !satisfied {
			 return false, nil
		}
	}

	return true, nil
}