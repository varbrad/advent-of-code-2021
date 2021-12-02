package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToIntegerHandlesABasicCase(t *testing.T) {
	expected := 2491
	actual := ToInteger("2491")

	assert.Equal(t, expected, actual)
}

func TestToIntegerReturnsAnError(t *testing.T) {
	assert.Panics(t, func() { ToInteger("2491a") })
}

func TestToIntegersHandlesABasicCase(t *testing.T) {
	expected := []int{1, 2, 3}
	actual := ToIntegers([]string{"1", "2", "3"})

	assert.Equal(t, expected, actual)
}

func TestToIntegersReturnsAnyErrors(t *testing.T) {
	assert.Panics(t, func() { ToIntegers([]string{"1", "b", "3"}) })
}
