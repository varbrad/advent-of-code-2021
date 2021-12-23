package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetXDiff(t *testing.T) {
	t.Run("Should solve an example 1", func(t *testing.T) {
		hole := location{variant: "hole", index: 0, value: 0, holePos: 0}
		rest := location{variant: "rest", index: 0, value: 0}

		expected := 2
		actual := getXDiff(hole, rest)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve an example 2", func(t *testing.T) {
		hole := location{variant: "hole", index: 0, value: 0, holePos: 0}
		rest := location{variant: "rest", index: 1, value: 0}

		expected := 1
		actual := getXDiff(hole, rest)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve an example 3", func(t *testing.T) {
		hole := location{variant: "hole", index: 0, value: 0, holePos: 0}
		rest := location{variant: "rest", index: 2, value: 0}

		expected := 1
		actual := getXDiff(hole, rest)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve an example 4", func(t *testing.T) {
		hole := location{variant: "hole", index: 0, value: 0, holePos: 0}
		rest := location{variant: "rest", index: 3, value: 0}

		expected := 3
		actual := getXDiff(hole, rest)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve an example 4", func(t *testing.T) {
		hole := location{variant: "hole", index: 0, value: 0, holePos: 0}
		rest := location{variant: "rest", index: 4, value: 0}

		expected := 5
		actual := getXDiff(hole, rest)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve an example 4", func(t *testing.T) {
		hole := location{variant: "hole", index: 0, value: 0, holePos: 0}
		rest := location{variant: "rest", index: 5, value: 0}

		expected := 7
		actual := getXDiff(hole, rest)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve an example 4", func(t *testing.T) {
		hole := location{variant: "hole", index: 0, value: 0, holePos: 0}
		rest := location{variant: "rest", index: 6, value: 0}

		expected := 8
		actual := getXDiff(hole, rest)

		assert.Equal(t, expected, actual)
	})
}
