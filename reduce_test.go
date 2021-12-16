package satsolver

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReduce(t *testing.T) {
	t.Run("a", func(t *testing.T) {
		f := MustParse("a")
		vars := VarsMap(f)

		f, err := reduce(f)
		require.NoError(t, err)

		require.Len(t, f.C, 0)
		require.True(t, vars["a"].Value)
		require.True(t, vars["a"].Assigned)
	})

	t.Run("~a", func(t *testing.T) {
		f := MustParse("~a")
		vars := VarsMap(f)

		f, err := reduce(f)
		require.NoError(t, err)

		require.Len(t, f.C, 0)
		require.False(t, vars["a"].Value)
		require.True(t, vars["a"].Assigned)
	})

	t.Run("unit rule", func(t *testing.T) {
		f := MustParse("a ^ (a v b)")
		vars := VarsMap(f)

		f, err := reduce(f)
		require.NoError(t, err)

		require.Len(t, f.C, 0)
		require.True(t, vars["a"].Value)
		require.True(t, vars["a"].Assigned)
		require.False(t, vars["b"].Assigned)
	})

	t.Run("tautology", func(t *testing.T) {
		f := MustParse("(a v ~a) ^ (a v b)")
		vars := VarsMap(f)
		f, err := reduce(f)
		require.NoError(t, err)

		require.Len(t, f.C, 1)
		require.Len(t, f.C[0].L, 2)

		require.False(t, f.C[0].L[0].Negated)
		require.Equal(t, "a", f.C[0].L[0].V.Name)
		require.False(t, vars["a"].Assigned)

		require.False(t, f.C[0].L[1].Negated)
		require.Equal(t, "b", f.C[0].L[1].V.Name)
		require.False(t, vars["b"].Assigned)
	})

	t.Run("contradiction", func(t *testing.T) {
		f := MustParse("a ^ ~a")
		_, err := reduce(f)
		require.Error(t, err)
	})

	t.Run("(a v b v c) ^ (~a v ~b v c) ^ (a v ~b v ~c) ^ (~a v b v ~c)", func(t *testing.T) {
		f := MustParse("(a v b v c) ^ (~a v ~b v c) ^ (a v ~b v ~c) ^ (~a v b v ~c)")
		vars := VarsMap(f)
		require.Len(t, f.C, 4)

		// nothing gets reduced
		f2, err := reduce(f)
		require.NoError(t, err)
		require.Len(t, f2.C, 4)
		require.Equal(t, f.String(), f2.String())

		// add a=true
		f2.C = append(f2.C, Clause{
			L: []Literal{
				{
					Negated: false,
					V:       vars["a"],
				},
			},
		})
		f3, err := reduce(f2)
		require.NoError(t, err)
		require.Len(t, f3.C, 2)
		require.True(t, vars["a"].Value)
		require.True(t, vars["a"].Assigned)
		require.False(t, vars["b"].Assigned)
		require.False(t, vars["c"].Assigned)

		// add b=true
		f3.C = append(f3.C, Clause{
			L: []Literal{
				{
					Negated: false,
					V:       vars["b"],
				},
			},
		})
		f4, err := reduce(f3)
		require.NoError(t, err)
		require.Len(t, f4.C, 1)
		require.True(t, vars["a"].Value)
		require.True(t, vars["a"].Assigned)
		require.True(t, vars["b"].Value)
		require.True(t, vars["b"].Assigned)
		require.False(t, vars["c"].Assigned)

		// add c=true
		f4.C = append(f4.C, Clause{
			L: []Literal{
				{
					Negated: false,
					V:       vars["c"],
				},
			},
		})
		f5, err := reduce(f4)
		require.NoError(t, err)
		require.Len(t, f5.C, 0)
		require.True(t, vars["a"].Value)
		require.True(t, vars["a"].Assigned)
		require.True(t, vars["b"].Value)
		require.True(t, vars["b"].Assigned)
		require.True(t, vars["c"].Value)
		require.True(t, vars["c"].Assigned)
	})
}