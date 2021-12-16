package satsolver

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenerate3SAT(t *testing.T) {
	t.Run("m=1,n=10", func(t *testing.T) {
		f, err := GenerateUniform3SAT(1, 10, nil)
		require.NoError(t, err)
		require.Len(t, f.C, 1)
	})

	t.Run("m=3,n=9", func(t *testing.T) {
		f, err := GenerateUniform3SAT(3, 9, nil)
		require.NoError(t, err)
		require.Len(t, f.C, 3)
	})

	t.Run("always the same", func(t *testing.T) {
		m := 8
		n := 2
		seed := int64(12345)
		first, err := GenerateUniform3SAT(m, n, &seed)
		require.NoError(t, err)

		for i := 0; i < 50; i++ {
			f, err := GenerateUniform3SAT(m, n, &seed)
			require.NoError(t, err)
			require.Equal(t, first.String(), f.String())
		}
	})
}
