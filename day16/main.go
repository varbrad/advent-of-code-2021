package main

import (
	"strconv"

	"github.com/varbrad/advent-of-code-2021/utils"
)

func main() {
	input := utils.ReadInputToString("day16/input")
	utils.Day(16, "Packet Decoder").Input(input).Part1(Day16Part1).Part2(Day16Part2)
}

func Day16Part1(input string) int {
	t := parseToTransmission(input)
	packet, _ := t.makePacket(0)

	return packet.versionSum()
}

func Day16Part2(input string) int {
	t := parseToTransmission(input)
	packet, _ := t.makePacket(0)

	return packet.value()
}

func parseToTransmission(input string) *transmission {
	bits := make([]int, len(input)*4)
	for i, c := range input {
		val, ok := strconv.ParseInt(string(c), 16, 0)
		if ok != nil {
			panic("invalid input")
		}
		for x := 3; x >= 0; x-- {
			bits[i*4+x] = int(val & 0x1)
			val = val >> 1
		}
	}
	return &transmission{bits: bits}
}

type transmission struct {
	bits []int
}

func (t *transmission) makePacket(index int) (*packet, int) {
	p := &packet{}
	p.version = bitSliceToInt(&t.bits, index, 3)
	p.typeId = bitSliceToInt(&t.bits, index+3, 3)

	parsedBits := 6

	// literal
	if p.typeId == 4 {
		literal := []int{}
		for {
			ni := index + parsedBits
			lb := t.bits[ni]
			literal = append(literal, t.bits[ni+1:ni+1+4]...)
			parsedBits += 5
			if lb == 0 {
				break
			}
		}
		p.literal = bitSliceToInt(&literal, 0, len(literal))
	} else {
		p.lengthTypeId = bitSliceToInt(&t.bits, index+parsedBits, 1)
		parsedBits++

		subPackets := []*packet{}
		if p.lengthTypeId == 0 {
			len := bitSliceToInt(&t.bits, index+parsedBits, 15)
			parsedBits += 15
			a := 0
			for a < len {
				sp, tb := t.makePacket(index + parsedBits)
				parsedBits += tb
				a += tb
				subPackets = append(subPackets, sp)
			}
		} else if p.lengthTypeId == 1 {
			pkts := bitSliceToInt(&t.bits, index+parsedBits, 11)
			parsedBits += 11
			for i := 0; i < pkts; i++ {
				sp, tb := t.makePacket(index + parsedBits)
				parsedBits += tb
				subPackets = append(subPackets, sp)
			}
		}
		p.subPackets = subPackets
	}

	return p, totalParsedBits(parsedBits)
}

func bitSliceToInt(bits *[]int, i int, l int) int {
	val := 0
	sl := (*bits)[i : i+l]
	for i, b := range sl {
		val += b << int(l-i-1)
	}
	return val
}

func totalParsedBits(parsedBits int) int {
	return parsedBits
}
