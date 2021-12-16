package main

import "math"

type packet struct {
	version      int
	typeId       int
	literal      int
	lengthTypeId int
	subPackets   []*packet
}

func (p *packet) versionSum() int {
	if p.subPackets == nil || len(p.subPackets) == 0 {
		return p.version
	}
	sum := p.version
	for _, sp := range p.subPackets {
		sum += sp.versionSum()
	}
	return sum
}

func (p *packet) value() int {
	switch p.typeId {
	case 0:
		return p.sum()
	case 1:
		return p.product()
	case 2:
		return p.min()
	case 3:
		return p.max()
	case 4:
		return p.literal
	case 5:
		return p.gt()
	case 6:
		return p.lt()
	case 7:
		return p.eq()
	}
	panic("unknown typeId found on a packet")
}

func (p *packet) sum() int {
	sum := 0
	for _, sp := range p.subPackets {
		sum += sp.value()
	}
	return sum
}

func (p *packet) product() int {
	prod := 1
	for _, sp := range p.subPackets {
		prod *= sp.value()
	}
	return prod
}

func (p *packet) min() int {
	min := math.MaxInt
	for _, sp := range p.subPackets {
		v := sp.value()
		if v < min {
			min = v
		}
	}
	return min
}

func (p *packet) max() int {
	max := math.MinInt
	for _, sp := range p.subPackets {
		v := sp.value()
		if v > max {
			max = v
		}
	}
	return max
}

func (p *packet) gt() int {
	a := p.subPackets[0].value()
	b := p.subPackets[1].value()
	if a > b {
		return 1
	}
	return 0
}

func (p *packet) lt() int {
	a := p.subPackets[0].value()
	b := p.subPackets[1].value()
	if a < b {
		return 1
	}
	return 0
}

func (p *packet) eq() int {
	a := p.subPackets[0].value()
	b := p.subPackets[1].value()
	if a == b {
		return 1
	}
	return 0
}
