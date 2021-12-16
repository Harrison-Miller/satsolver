package satsolver

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func testVerify(t *testing.T, formula string, I Interpretation, expected bool) {
	f := MustParse(formula)
	satisfied, err := Verify(f, I)
	require.NoError(t, err)
	require.Equal(t, expected, satisfied)
}

func TestVerify(t *testing.T) {
	t.Run("a, a=true", func(t *testing.T) {
		testVerify(t, "a", Interpretation{"a": true}, true)
	})

	t.Run("a, a=false", func(t *testing.T) {
		testVerify(t, "a", Interpretation{"a": false}, false)
	})

	t.Run("~a, a=true", func(t *testing.T) {
		testVerify(t, "~a", Interpretation{"a": true}, false)
	})

	t.Run("~a, a=false", func(t *testing.T) {
		testVerify(t, "~a", Interpretation{"a": false}, true)
	})

	t.Run("a ^ ~a", func(t *testing.T) {
		testVerify(t, "a ^ ~a", Interpretation{"a": true}, false)
		testVerify(t, "a ^ ~a", Interpretation{"a": false}, false)
	})

	t.Run("(a v ~a)", func(t *testing.T) {
		testVerify(t, "(a v ~a)", Interpretation{"a": true}, true)
		testVerify(t, "(a v ~a)", Interpretation{"a": false}, true)
	})

	t.Run("a ^ b", func(t *testing.T) {
		testVerify(t, "a ^ b", Interpretation{"a": true, "b":true}, true)
		testVerify(t, "a ^ b", Interpretation{"a": true, "b":false}, false)
		testVerify(t, "a ^ b", Interpretation{"a": false, "b":true}, false)
		testVerify(t, "a ^ b", Interpretation{"a": false, "b":false}, false)
	})

	t.Run("(a v b)", func(t *testing.T) {
		testVerify(t, "(a v b)", Interpretation{"a": true, "b":true}, true)
		testVerify(t, "(a v b)", Interpretation{"a": true, "b":false}, true)
		testVerify(t, "(a v b)", Interpretation{"a": false, "b":true}, true)
		testVerify(t, "(a v b)", Interpretation{"a": false, "b":false}, false)
	})

	t.Run("a ^ (~b v c) ^ d", func(t *testing.T) {
		testVerify(t, "a ^ (~b v c) ^ d", Interpretation{"a": true, "b":true, "c":true, "d":true}, true)
		testVerify(t, "a ^ (~b v c) ^ d", Interpretation{"a": false, "b":true, "c":true, "d":true}, false)
		testVerify(t, "a ^ (~b v c) ^ d", Interpretation{"a": true, "b":true, "c":true, "d":false}, false)
		testVerify(t, "a ^ (~b v c) ^ d", Interpretation{"a": true, "b":true, "c":false, "d":true}, false)
		testVerify(t, "a ^ (~b v c) ^ d", Interpretation{"a": true, "b":false, "c":false, "d":true}, true)
	})
}