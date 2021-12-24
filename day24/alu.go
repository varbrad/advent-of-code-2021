package main

import (
	"fmt"
	"strconv"
	"strings"
)

type op struct {
	cmd  string
	arg0 rune
	arg1 rune
	num1 int
}
type program struct {
	index int
	ops   []op
	mem   map[rune]int
	input []int
}

func parseProgram(source []string) *program {
	ops := make([]op, len(source))
	for ix, line := range source {
		vs := strings.Split(line, " ")
		ops[ix] = op{
			cmd:  vs[0],
			arg0: rune(vs[1][0]),
		}
		if len(vs) > 2 {
			v, err := strconv.Atoi(vs[2])
			if err == nil {
				ops[ix].num1 = v
			} else {
				ops[ix].arg1 = rune(vs[2][0])
			}
		}
	}
	return &program{0, ops, map[rune]int{'w': 0, 'x': 0, 'y': 0, 'z': 0}, []int{}}
}

func (p *program) setInput(input []int) {
	p.input = input
}

func (p *program) run() {
	for p.index < len(p.ops) {
		p.step()
	}
}

func (p *program) step() {
	op := p.ops[p.index]

	switch op.cmd {
	case "inp":
		{
			val := p.input[0]
			p.input = p.input[1:]
			p.mem[op.arg0] = val
		}
	case "add":
		{
			v := op.num1
			if op.arg1 != 0 {
				v = p.mem[op.arg1]
			}
			p.mem[op.arg0] += v
		}
	case "mul":
		{
			v := op.num1
			if op.arg1 != 0 {
				v = p.mem[op.arg1]
			}
			p.mem[op.arg0] *= v
		}
	case "div":
		{
			v := op.num1
			if op.arg1 != 0 {
				v = p.mem[op.arg1]
			}
			p.mem[op.arg0] /= v
		}
	case "mod":
		{
			v := op.num1
			if op.arg1 != 0 {
				v = p.mem[op.arg1]
			}
			p.mem[op.arg0] %= v
		}
	case "eql":
		{
			b := op.num1
			if op.arg1 != 0 {
				b = p.mem[op.arg1]
			}
			a := p.mem[op.arg0]
			if a == b {
				p.mem[op.arg0] = 1
			} else {
				p.mem[op.arg0] = 0
			}
		}
	}

	p.index++
}

func (p *program) getMemory(r rune) int {
	return p.mem[r]
}

func (p *program) reset() {
	p.index = 0
	p.mem = map[rune]int{'w': 0, 'x': 0, 'y': 0, 'z': 0}
	p.input = []int{}
}

func (p *program) String() string {
	s := ""
	for _, op := range p.ops {
		str := ""
		if op.arg1 != 0 {
			str = string(op.arg1)
		} else {
			str = fmt.Sprint(op.num1)
		}
		s += fmt.Sprintf("%s %s %s\n", op.cmd, string(op.arg0), str)
	}
	s += "\n"
	for k, v := range p.mem {
		s += string(k) + " = " + fmt.Sprint(v) + "\n"
	}
	return s
}
