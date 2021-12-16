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
}