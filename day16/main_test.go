package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/varbrad/advent-of-code-2021/utils"
)

func TestDay16Part1(t *testing.T) {
	t.Run("Should solve an example #0", func(t *testing.T) {
		input := "D2FE28"

		expected := 6
		actual := Day16Part1(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve an example #1", func(t *testing.T) {
		input := "38006F45291200"

		expected := 9
		actual := Day16Part1(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve an example #2", func(t *testing.T) {
		input := "EE00D40C823060"

		expected := 14
		actual := Day16Part1(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve an example #3", func(t *testing.T) {
		input := "8A004A801A8002F478"

		expected := 16
		actual := Day16Part1(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve an example #4", func(t *testing.T) {
		input := "620080001611562C8802118E34"

		expected := 12
		actual := Day16Part1(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve an example #5", func(t *testing.T) {
		input := "C0015000016115A2E0802F182340"

		expected := 23
		actual := Day16Part1(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve an example #6", func(t *testing.T) {
		input := "A0016C880162017C3686B18A3D4780"

		expected := 31
		actual := Day16Part1(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve my puzzle input", func(t *testing.T) {
		input := utils.ReadInputToString("day16/input")

		expected := 897
		actual := Day16Part1(input)

		assert.Equal(t, expected, actual)
	})
}

func TestDay16Part2(t *testing.T) {
	t.Run("Should solve an example #0", func(t *testing.T) {
		input := "C200B40A82"

		expected := 3
		actual := Day16Part2(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve an example #1", func(t *testing.T) {
		input := "04005AC33890"

		expected := 54
		actual := Day16Part2(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve an example #2", func(t *testing.T) {
		input := "880086C3E88112"

		expected := 7
		actual := Day16Part2(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve an example #3", func(t *testing.T) {
		input := "CE00C43D881120"

		expected := 9
		actual := Day16Part2(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve an example #4", func(t *testing.T) {
		input := "D8005AC2A8F0"

		expected := 1
		actual := Day16Part2(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve an example #5", func(t *testing.T) {
		input := "F600BC2D8F"

		expected := 0
		actual := Day16Part2(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve an example #6", func(t *testing.T) {
		input := "9C005AC2F8F0"

		expected := 0
		actual := Day16Part2(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve an example #7", func(t *testing.T) {
		input := "9C0141080250320F1802104A08"

		expected := 1
		actual := Day16Part2(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve my puzzle input", func(t *testing.T) {
		input := utils.ReadInputToString("day16/input")

		expected := 9485076995911
		actual := Day16Part2(input)

		assert.Equal(t, expected, actual)
	})
}
