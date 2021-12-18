package main

import (
	"fmt"
	"regexp"

	"github.com/varbrad/advent-of-code-2021/utils"
)

type snailfishNumber struct {
	value int
	left  *snailfishNumber
	right *snailfishNumber
}

var numberRegex = regexp.MustCompile(`^\d+$`)

func parseSnailfishNumber(input string) *snailfishNumber {
	return getNextSnailfish(input)
}

func getNextSnailfish(input string) *snailfishNumber {
	if numberRegex.MatchString(input) {
		return &snailfishNumber{value: utils.ToInteger(input)}
	}

	stack := 0
	midPoint := -1
	for i, r := range input {
		switch {
		case r == ',' && stack == 1:
			midPoint = i
			break
		case r == '[':
			stack++
		case r == ']':
			stack--
		}
	}

	if midPoint == -1 {
		panic("Unable to parse input")
	}

	newSfn := &snailfishNumber{}
	newSfn.left = getNextSnailfish(input[1:midPoint])
	newSfn.right = getNextSnailfish(input[midPoint+1 : len(input)-1])
	return newSfn
}

// methods

func (sfn *snailfishNumber) String() string {
	if sfn.left == nil && sfn.right == nil {
		return fmt.Sprint(sfn.value)
	}
	return "[" + sfn.left.String() + "," + sfn.right.String() + "]"
}

func (sfn *snailfishNumber) Add(sfn2 *snailfishNumber) *snailfishNumber {
	return &snailfishNumber{
		left:  sfn,
		right: sfn2,
	}
}

func (sfn *snailfishNumber) Reduce() *snailfishNumber {
	reduceNode := sfn.findExplodeNode(0)
	if reduceNode != nil {
		sfn.explode(reduceNode)
		return sfn.Reduce()
	}

	splitNode := sfn.findSplitNode()
	if splitNode != nil {
		sfn.split(splitNode)
		return sfn.Reduce()
	}

	return sfn
}

func (sfn *snailfishNumber) findExplodeNode(depth int) *snailfishNumber {
	if depth == 4 && sfn.left != nil && sfn.right != nil {
		return sfn
	}

	if sfn.left != nil {
		v := sfn.left.findExplodeNode(depth + 1)
		if v != nil {
			return v
		}
	}

	if sfn.right != nil {
		v := sfn.right.findExplodeNode(depth + 1)
		if v != nil {
			return v
		}
	}

	return nil
}

func (sfn *snailfishNumber) explode(explodeNode *snailfishNumber) {
	values := sfn.list()
	indexL := explodeNode.left.index(values)

	if indexL > 0 {
		values[indexL-1].value += explodeNode.left.value
	}
	if indexL+2 < len(values) {
		values[indexL+2].value += explodeNode.right.value
	}

	explodeNode.left = nil
	explodeNode.right = nil
	explodeNode.value = 0
}

func (sfn *snailfishNumber) findSplitNode() *snailfishNumber {
	if sfn.left == nil && sfn.right == nil {
		if sfn.value > 9 {
			return sfn
		}
		return nil
	}

	if sfn.left != nil {
		v := sfn.left.findSplitNode()
		if v != nil {
			return v
		}
	}

	if sfn.right != nil {
		v := sfn.right.findSplitNode()
		if v != nil {
			return v
		}
	}

	return nil
}

func (sfn *snailfishNumber) split(splitNode *snailfishNumber) {
	left := splitNode.value / 2
	right := (splitNode.value + 1) / 2

	splitNode.value = 0
	splitNode.left = &snailfishNumber{value: left}
	splitNode.right = &snailfishNumber{value: right}
}

func (sfn *snailfishNumber) list() []*snailfishNumber {
	if sfn.left == nil && sfn.right == nil {
		return []*snailfishNumber{sfn}
	}

	left := sfn.left.list()
	right := sfn.right.list()

	return append(left, right...)
}

func (sfn *snailfishNumber) index(list []*snailfishNumber) int {
	for i, n := range list {
		if n == sfn {
			return i
		}
	}
	return -1
}

func (sfn *snailfishNumber) Magnitude() int {
	if sfn.left == nil && sfn.right == nil {
		return sfn.value
	}
	return sfn.left.Magnitude()*3 + sfn.right.Magnitude()*2
}
