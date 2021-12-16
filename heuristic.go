package satsolver

import (
	"math"
	"math/rand"
)

func getUnassigned(f Formula) []*Variable {
	variables := make([]*Variable, 0)
	for _, v := range f.V {
		if !v.Assigned {
			variables = append(variables, v)
		}
	}
	return variables
}

func FirstUnassigned(f Formula) *Variable {
	// TODO: This is actually random because maps don't always iterate in the same order
	v := getUnassigned(f)
	if len(v) > 0 {
		return v[0]
	}
	return nil
}

func RandomUnassigned(f Formula) *Variable {
	v := getUnassigned(f)
	if len(v) > 0 {
		return v[rand.Intn(len(v))]
	}
	return nil
}

func MaximumOccurrences(f Formula) *Variable {
	var maxV *Variable
	var count int
	for _, v := range f.V {
		current := 0
		if !v.Assigned {
			// find the number of times it occurs
			for _, clause := range f.C {
				for _, literal := range clause.L {
					if literal.V == v {
						current++
					}
				}
			}

			if current > count {
				count = current
				maxV = v
			}
		}
	}

	return maxV
}

func JeroslowWang(f Formula) *Variable {
	weights := map[*Variable]float64{}
	for _, clause := range f.C {
		for _, literal := range clause.L {
			if !literal.V.Assigned {
				w := weights[literal.V]
				w += math.Pow(2, -float64(len(clause.L)))
				weights[literal.V] = w
			}
		}
	}

	var maxV *Variable
	var max float64
	for v, w := range weights {
		if w > max {
			max = w
			maxV = v
		}
	}

	return maxV
}