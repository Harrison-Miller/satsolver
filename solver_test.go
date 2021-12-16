package satsolver

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func testSolver(t *testing.T, formula string, expected bool) {
	f := MustParse(formula)
	I, err := Solve(f, FirstUnassigned)
	if !expected {
		require.Error(t, err)
	} else {
		require.NoError(t, err)
		satisfied, err := Verify(f, I)
		require.NoError(t, err)
		require.Equal(t, expected, satisfied, f.String())
	}
}

func TestSolver(t *testing.T) {
	t.Run("a", func(t *testing.T) {testSolver(t, "a", true)})
	t.Run("~a", func(t *testing.T) {testSolver(t, "~a", true)})
	t.Run("a ^ b", func(t *testing.T) {testSolver(t, "a ^ b", true)})
	t.Run("a ^ ~a", func(t *testing.T) {testSolver(t, "a ^ ~a", false)})
	t.Run("(a v ~a)", func(t *testing.T) {testSolver(t, "(a v ~a)", true)})
	t.Run("(a v b)", func(t *testing.T) {testSolver(t, "(a v b)", true)})
	t.Run("a ^ (~b v c) ^ d", func(t *testing.T) {testSolver(t, "a ^ (~b v c) ^ d", true)})
	t.Run("(x1 v ~x2) ^ (~x1 v x2 v x3) ^ ~x1", func(t *testing.T) {testSolver(t, "(x1 v ~x2) ^ (~x1 v x2 v x3) ^ ~x1", true)})
}

func TestComplex(t *testing.T) {
	t.Run("(a v b v c) ^ (~a v ~b v c) ^ (a v ~b v ~c) ^ (~a v b v ~c)", func(t *testing.T) {
		testSolver(t, "(a v b v c) ^ (~a v ~b v c) ^ (a v ~b v ~c) ^ (~a v b v ~c)", true)
	})
	
	t.Run("(~a v ~b v ~c) ^ (~a v ~b v c) ^ (a v ~b v ~c) ^ (~a v b v ~c)", func(t *testing.T) {
		testSolver(t, "(~a v ~b v ~c) ^ (~a v ~b v c) ^ (a v ~b v ~c) ^ (~a v b v ~c)", true)
	})

	t.Run("big", func(t *testing.T) {
		s := int64(123456)
		f, err := GenerateUniform3SAT(100, 23, &s)
		require.NoError(t, err)
		I, err := Solve(f, JeroslowWang)
		require.NoError(t, err)
		satisfied, err := Verify(f, I)
		require.NoError(t, err)
		require.True(t, satisfied)
	})
}

func benchmarkSolverWith3SAT(b *testing.B, m int, n int, seed int64) {
	f, err := GenerateUniform3SAT(m, n, &seed)
	//fmt.Println(f)
	require.NoError(b, err)

	// make sure this is a solvable instance
	I, err := Solve(f, JeroslowWang)
	require.NoError(b, err)
	satisfied, err := Verify(f, I)
	require.NoError(b, err)
	require.True(b, satisfied)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Solve(f, JeroslowWang)
	}
}

func BenchmarkSolverM8N2(b *testing.B) { benchmarkSolverWith3SAT(b, 8, 2, 12345)}
func BenchmarkSolverM17N4(b *testing.B) { benchmarkSolverWith3SAT(b, 17, 4, 12345)}
func BenchmarkSolverM39N9(b *testing.B) { benchmarkSolverWith3SAT(b, 39, 9, 12345)}
func BenchmarkSolverM50N12(b *testing.B) { benchmarkSolverWith3SAT(b, 50, 12, 12345)}
func BenchmarkSolverM80N50(b *testing.B) { benchmarkSolverWith3SAT(b, 80, 50, 123456)}
func BenchmarkSolverM100N23(b *testing.B) { benchmarkSolverWith3SAT(b, 100, 23, 123456)}
func BenchmarkSolverM150N35(b *testing.B) { benchmarkSolverWith3SAT(b, 150, 35, 123456)}
func BenchmarkSolverM200N47(b *testing.B) { benchmarkSolverWith3SAT(b, 200, 47, 123456)}