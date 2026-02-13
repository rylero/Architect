package library

import (
	"architect/sim"
	"fmt"
)

/* Basic Composite Gates */

func (b *Builder) NAND(inA, inB sim.NetID, width uint8) sim.NetID {
	andOut := b.AllocNet(width)
	nandOut := b.AllocNet(width)

	andNode := &AND{InA: inA, InB: inB, Out: andOut}
	b.AddNode(andNode)

	notNode := &NOT{In: andOut, Out: nandOut}
	b.AddNode(notNode)

	return nandOut
}

func (b *Builder) NOR(inA, inB sim.NetID, width uint8) sim.NetID {
	orOut := b.AllocNet(width)
	nandOut := b.AllocNet(width)

	orNode := &OR{InA: inA, InB: inB, Out: orOut}
	b.AddNode(orNode)

	notNode := &NOT{In: orOut, Out: nandOut}
	b.AddNode(notNode)
	return nandOut
}

/* Adder */
func (b *Builder) HalfAdder(inA, inB sim.NetID, width uint8) (sum, carry sim.NetID) {
	b.EnterScope("halfAdder")
	defer b.ExitScope()

	sum = b.AllocNamedNet(width, "sum")
	carry = b.AllocNamedNet(width, "carry")

	xorNode := &XOR{InA: inA, InB: inB, Out: sum}
	b.AddNode(xorNode)

	andNode := &AND{InA: inA, InB: inB, Out: carry}
	b.AddNode(andNode)

	return sum, carry
}

func (b *Builder) FullAdder(inA, inB, inC sim.NetID, width uint8) (sum, carry sim.NetID) {
	b.EnterScope("fullAdder")
	defer b.ExitScope()

	abSum, abCarry := b.HalfAdder(inA, inB, width)
	ccSum, ccCarry := b.HalfAdder(abSum, inC, width)

	carry = b.AllocNamedNet(width, "carry")
	orNode := &OR{InA: abCarry, InB: ccCarry, Out: carry}
	b.AddNode(orNode)

	return ccSum, carry
}

func (b *Builder) RippleCarryAdder(inA, inB, cin sim.NetID, width uint8) (sum, carry sim.NetID) {
	b.EnterScope("ripple")
	defer b.ExitScope()

	aBits := b.Split(inA, width)
	bBits := b.Split(inB, width)
	carry = cin

	outputs := make([]sim.NetID, width)

	for i := uint8(0); i < width; i++ {
		b.EnterScope(fmt.Sprintf("%d", i))
		faSum, faCarry := b.FullAdder(aBits[i], bBits[i], carry, 1)
		b.ExitScope()

		carry = faCarry
		outputs[i] = faSum
	}

	sum = b.Join(outputs, width)

	return sum, carry
}

/* Memory */
func (b *Builder) SRLatchNOR(set, reset sim.NetID) (sim.NetID, sim.NetID) {
	width := uint8(1)

	q := b.AllocNamedNet(width, "Q")
	qBar := b.AllocNamedNet(width, "QBar")

	nor1Out := b.NOR(set, qBar, width)
	nor2Out := b.NOR(reset, q, width)

	b.WireInto(nor1Out, q)
	b.WireInto(nor2Out, qBar)

	return q, qBar
}

func (b *Builder) DLatch(d, en sim.NetID) (sim.NetID, sim.NetID) {
	width := uint8(1)

	dNot := b.NOT(d, width)
	s := b.AND(en, d, width)
	r := b.AND(en, dNot, width)

	return b.SRLatchNOR(s, r)
}
