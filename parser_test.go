package satsolver

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	t.Run("a", func(t *testing.T) {
		f, err := Parse(strings.NewReader("a"))
		require.NoError(t, err)

		require.Len(t, f.C, 1)
		require.Len(t, f.C[0].L, 1)
		require.False(t, f.C[0].L[0].Negated)
		require.Equal(t, "a", f.C[0].L[0].V.Name)
	})

	t.Run("~a", func(t *testing.T) {
		f, err := Parse(strings.NewReader("~a"))
		require.NoError(t, err)

		require.Len(t, f.C, 1)
		require.Len(t, f.C[0].L, 1)
		require.True(t, f.C[0].L[0].Negated)
		require.Equal(t, "a", f.C[0].L[0].V.Name)
	})

	t.Run("a ^ b", func(t *testing.T) {
		f, err := Parse(strings.NewReader("a ^ ~b"))
		require.NoError(t, err)

		require.Len(t, f.C, 2)
		require.Len(t, f.C[0].L, 1)
		require.Len(t, f.C[1].L, 1)

		require.False(t, f.C[0].L[0].Negated)
		require.Equal(t, "a", f.C[0].L[0].V.Name)

		require.True(t, f.C[1].L[0].Negated)
		require.Equal(t, "b", f.C[1].L[0].V.Name)
	})

	t.Run("(a v b)", func(t *testing.T) {
		f, err := Parse(strings.NewReader("(a v b)"))
		require.NoError(t, err)

		require.Len(t, f.C, 1)
		require.Len(t, f.C[0].L, 2)

		require.False(t, f.C[0].L[0].Negated)
		require.Equal(t, "a", f.C[0].L[0].V.Name)

		require.False(t, f.C[0].L[1].Negated)
		require.Equal(t, "b", f.C[0].L[1].V.Name)
	})

	t.Run("a ^ (~b v c) ^ d", func(t *testing.T) {
		f, err := Parse(strings.NewReader("a ^ (~b v c) ^ d"))
		require.NoError(t, err)

		require.Len(t, f.C, 3)
		require.Len(t, f.C[0].L, 1)
		require.Len(t, f.C[1].L, 2)
		require.Len(t, f.C[2].L, 1)

		require.False(t, f.C[0].L[0].Negated)
		require.Equal(t, "a", f.C[0].L[0].V.Name)

		require.True(t, f.C[1].L[0].Negated)
		require.Equal(t, "b", f.C[1].L[0].V.Name)

		require.False(t, f.C[1].L[1].Negated)
		require.Equal(t, "c", f.C[1].L[1].V.Name)

		require.False(t, f.C[2].L[0].Negated)
		require.Equal(t, "d", f.C[2].L[0].V.Name)
	})
	
	t.Run("a ^ a", func(t *testing.T) {
		f, err := Parse(strings.NewReader("a ^ a"))
		require.NoError(t, err)

		require.Len(t, f.C, 2)
		require.Len(t, f.C[0].L, 1)
		require.Len(t, f.C[1].L, 1)

		require.False(t, f.C[0].L[0].Negated)
		require.Equal(t, "a", f.C[0].L[0].V.Name)

		require.False(t, f.C[1].L[0].Negated)
		require.Equal(t, "a", f.C[1].L[0].V.Name)

		require.Same(t, f.C[0].L[0].V, f.C[1].L[0].V)
	})

}
