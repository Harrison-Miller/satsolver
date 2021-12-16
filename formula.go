package satsolver

import (
	"fmt"
	"strings"
)

// CNF Grammar
// Formula --> Clause { "^" Clause}
// Clause --> Literal | "(" Literal { "v" Literal } ")"
// Literal --> ["~"]Variable
// Variable --> Ident

type Formula struct {
	C []Clause `@@ ("^" @@)*`
	V map[string]*Variable
}

func CloneFormula(f Formula) Formula {
	f2 := Formula{
		C:  make([]Clause, len(f.C)),
		V: f.V,
	}
	copy(f2.C, f.C)

	// copy each literal in the clause as well
	for i, clause := range f.C {
		f2.C[i].L = make([]Literal, len(clause.L))
		copy(f2.C[i].L, clause.L)
	}

	return f2
}

func (f Formula) String() string {
	clauseStrs := []string{}
	for _, clause := range f.C {
		clauseStrs = append(clauseStrs, clause.String())
	}
	return strings.Join(clauseStrs, " ^ ")
}

type Clause struct {
	L []Literal `@@ | "(" @@ ("v" @@ )* ")"`
}

func (c Clause) String() string {
	if len(c.L) == 1 {
		return c.L[0].String()
	}

	literalStrs := []string{}
	for _, literal := range c.L {
		literalStrs = append(literalStrs, literal.String())
	}
	return fmt.Sprintf("(%s)", strings.Join(literalStrs, " v "))
}

// strict equal including order
func (a Clause) Equal(b Clause) bool {
	if len(a.L) != len(b.L) {
		return false
	}

	for i, _ := range a.L {
		if a.L[i].Negated != b.L[i].Negated || a.L[i].V != b.L[i].V {
			return false
		}
	}
	return true
}

type Literal struct {
	Negated bool `@"~"?`
	V *Variable `@@`
}

func (l Literal) String() string {
	if l.Negated {
		return fmt.Sprintf("~%s", l.V.Name)
	}

	return l.V.Name
}

type Variable struct {
	Name string `@Ident`
	Value bool
	Assigned bool
}

func (v Variable) String() string {
	return fmt.Sprintf("%s=%t", v.Name, v.Value)
}

func (v *Variable) Assign(value bool) {
	v.Value = value
	v.Assigned = true
}

type Interpretation map[string]bool

func (i Interpretation) String() string {
	builder := strings.Builder{}
	for s, b := range i {
		builder.WriteString(fmt.Sprintf("%s = %t\n", s, b))
	}
	return builder.String()
}

// Vars returns all variables in the formula
func Vars(f Formula) []*Variable {
	varsMap := map[string]*Variable{}
	vars := []*Variable{}
	for _, clause := range f.C {
		for _, literal := range clause.L {
			if _, ok := varsMap[literal.V.Name]; !ok {
				varsMap[literal.V.Name] = literal.V
				vars = append(vars, literal.V)
			}
		}
	}
	return vars
}

// VarsMap returns all variables in the formula by their name
func VarsMap(f Formula) map[string]*Variable {
	varsMap := map[string]*Variable{}
	for _, clause := range f.C {
		for _, literal := range clause.L {
			varsMap[literal.V.Name] = literal.V
		}
	}
	return varsMap
}

func FromVariables(vars []*Variable) Interpretation {
	I := Interpretation{}
	for _, v := range vars {
		I[v.Name] = v.Value
	}
	return I
}

func FromVariablesMap(vars map[string]*Variable) Interpretation {
	I := Interpretation{}
	for _, v := range vars {
		I[v.Name] = v.Value
	}
	return I
}