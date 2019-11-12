//
//    Copyright 2019 Insolar Technologies
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
//

package jetid

import (
	"fmt"
	"io"
	"math"
	"math/bits"
	"strings"

	"github.com/insolar/insolar/longbits"
)

type Prefix uint32

const SplitMedian = 7 // makes 0 vs 1 ratio like [0..6] vs [7..15]
// this enables left branches of jets to be ~23% less loaded

// limited to 65k Jets
type PrefixTree struct {
	mask       Prefix
	minDepth   uint8
	maxDepth   uint8
	leafCounts [17]uint16
	lenNibles  [32768]uint8
}

func (p *PrefixTree) MaxDepth() uint8 {
	return p.maxDepth
}

func (p *PrefixTree) MinDepth() uint8 {
	return p.minDepth
}

func (p *PrefixTree) getPrefixLength(prefix uint16) (uint8, bool) {
	depth := p.lenNibles[prefix>>1]
	if prefix&1 != 0 {
		depth >>= 4
	} else {
		depth &= 0x0F
	}
	switch {
	case depth != 0:
		return depth + 1, true
	case p.maxDepth == 0:
		return 0, prefix == 0
	default:
		return 1, prefix <= 1
	}
}

func (p *PrefixTree) setPrefixLength(prefix uint16, depth uint8) {
	switch {
	case depth > 16:
		panic("illegal value")
	case depth != 0:
		depth--
	case prefix != 0:
		panic("illegal value")
	default:
		p.lenNibles[0] = 0
		return
	}

	idx := prefix >> 1
	d := p.lenNibles[idx]
	if prefix&1 != 0 {
		if d&0xF0 != 0 {
			panic("illegal state")
		}
		d = (d & 0x0F) | (depth << 4)
	} else {
		if d&0x0F != 0 {
			panic("illegal state")
		}
		d = (d & 0xF0) | (depth & 0x0F)
	}
	p.lenNibles[idx] = d
}

func (p *PrefixTree) resetPrefixLength(prefix uint16) {
	if prefix&1 != 0 {
		p.lenNibles[prefix>>1] &= 0x0F
	} else {
		p.lenNibles[prefix>>1] &= 0xF0
	}
}

func (p *PrefixTree) FindPrefixLength(prefix Prefix) uint8 {
	_, l := p.findPrefixLength(prefix)
	return l
}

func (p *PrefixTree) findPrefixLength(prefix Prefix) (uint16, uint8) {
	maskedPrefix := uint16(prefix & p.mask)

	if depth, ok := p.getPrefixLength(maskedPrefix); ok {
		return maskedPrefix, depth
	}

	for maskedPrefix > math.MaxUint8 {
		level := 7 + bits.Len8(uint8(maskedPrefix>>8))
		hiBit := uint16(1) << uint8(level)
		maskedPrefix ^= hiBit

		if depth, ok := p.getPrefixLength(maskedPrefix); ok {
			//p.setPrefixLength(uint16(prefix & p.mask), depth)
			return maskedPrefix, depth
		}
	}

	for maskedPrefix > 0 {
		level := bits.Len8(uint8(maskedPrefix)) - 1
		hiBit := uint16(1) << uint8(level)
		maskedPrefix ^= hiBit
		if depth, ok := p.getPrefixLength(maskedPrefix); ok {
			//p.setPrefixLength(uint16(prefix & p.mask), depth)
			return maskedPrefix, depth
		}
	}

	panic("illegal state")
}

func (p *PrefixTree) Split(prefix Prefix, prefixLimit uint8) {
	maskedPrefix, prefixLen := p.findPrefixLength(prefix)
	switch {
	case prefixLimit != prefixLen:
		panic("illegal value")
	case int(prefixLen) >= len(p.leafCounts):
		panic("illegal value") // TODO return as error?
	}

	switch n := p.leafCounts[prefixLen]; {
	case n > 1:
		p.leafCounts[prefixLen] = n - 1
	case n == 1:
		p.leafCounts[prefixLen] = 0
		if prefixLen == p.minDepth {
			p.minDepth++
		}
	case prefixLen == 0 && p.minDepth == 0 && p.maxDepth == 0:
		// zero state
		p.minDepth++
	default:
		panic("illegal state")
	}

	if prefixLen == p.maxDepth {
		p.maxDepth++
		p.mask = (p.mask << 1) | 1
	}

	pairedPrefix := maskedPrefix | (1 << prefixLen)
	prefixLen++
	p.resetPrefixLength(maskedPrefix)
	p.setPrefixLength(maskedPrefix, prefixLen)
	p.setPrefixLength(pairedPrefix, prefixLen)
	p.leafCounts[prefixLen] += 2
}

func (p *PrefixTree) Merge(prefix Prefix, prefixLimit uint8) {
	maskedPrefix, prefixLen := p.findPrefixLength(prefix)
	switch {
	case prefixLimit != prefixLen:
		panic("illegal value")
	case prefixLen == 0:
		panic("illegal value")
	}

	pairedPrefix := maskedPrefix | (1 << (prefixLen - 1))
	if maskedPrefix == pairedPrefix {
		panic("illegal value - only the zero-side is allowed to merge") // TODO return as error?
	}
	if _, pairedPrefixLen := p.findPrefixLength(Prefix(pairedPrefix)); pairedPrefixLen != prefixLen {
		panic("illegal value - unbalanced merge pair") // TODO return as error?
	}

	switch n := p.leafCounts[prefixLen]; {
	case n > 2:
		p.leafCounts[prefixLen] = n - 2
	case n == 2:
		switch p.maxDepth {
		case 0:
			panic("illegal state")
		case prefixLen:
			p.maxDepth--
			p.mask >>= 1
		}
		p.leafCounts[prefixLen] = 0
	default:
		panic("illegal state")
	}

	if prefixLen == p.minDepth {
		p.minDepth--
	}
	prefixLen--
	p.resetPrefixLength(pairedPrefix)
	p.resetPrefixLength(maskedPrefix)
	p.setPrefixLength(maskedPrefix, prefixLen)
	p.leafCounts[prefixLen]++
}

func (p *PrefixTree) String() string {
	return fmt.Sprintf("min=%d max=%d cnt=%v", p.minDepth, p.maxDepth, p.leafCounts)
}

func (p *PrefixTree) PrintTable() {
	fmt.Printf("min=%d max=%d cnt=%v\n", p.minDepth, p.maxDepth, p.leafCounts)
	for i := range p.lenNibles {
		prefix := Prefix(i << 1)
		if depth, ok := p.getPrefixLength(uint16(prefix)); ok && depth != 0 {
			fmt.Printf("%5d [%2d]", prefix, depth)
			p.printRow(prefix, depth)
		}

		prefix++
		if depth, ok := p.getPrefixLength(uint16(prefix)); ok && depth != 0 {
			fmt.Printf("%5d [%2d]", prefix, depth)
			p.printRow(prefix, depth)
		}
	}
}

func (p *PrefixTree) printRow(prefix Prefix, pLen uint8) {
	b := strings.Builder{}
	b.Grow(32)
	for i := uint8(0); i < pLen; i++ {
		b.WriteByte(' ')
		b.WriteByte('0' | byte(prefix)&1)
		prefix >>= 1
	}
	fmt.Println(b.String())
}

const simpleSerializeV1 = 0

// TODO Deserialize
func (p *PrefixTree) SimpleSerialize(w io.Writer) error {
	if p.maxDepth == 0 {
		panic("illegal state")
	}
	if _, e := w.Write([]byte{simpleSerializeV1, p.minDepth | p.maxDepth<<4}); e != nil {
		return e
	}
	delta := p.maxDepth - p.minDepth
	if delta == 0 {
		return nil
	}
	maxBound := 1 << (p.maxDepth - 1)
	switch encodingBits := bits.Len8(delta); encodingBits {
	case 4:
		_, e := w.Write(p.lenNibles[:maxBound])
		return e
	case 1:
		bb := longbits.NewBitBuilder(longbits.FirstLow, (maxBound+3)>>2)
		for _, b := range p.lenNibles[:maxBound] {
			bb.Append(b&0x0F > p.minDepth)
			bb.Append((b >> 4) > p.minDepth)
		}
		_, e := w.Write(bb.DoneToBytes())
		return e
	default:
		bb := longbits.NewBitBuilder(longbits.FirstLow, (maxBound+4-encodingBits)/encodingBits)
		for _, b := range p.lenNibles[:maxBound] {
			bb.AppendSubByte(b&0x0F-p.minDepth, uint8(encodingBits))
			bb.AppendSubByte((b>>4)-p.minDepth, uint8(encodingBits))
		}
		_, e := w.Write(bb.DoneToBytes())
		return e
	}
}