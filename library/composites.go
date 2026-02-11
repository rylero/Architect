package library

import "architect/sim"

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

func (b *Builder) HalfAdder(inA, inB sim.NetID) (sim.NetID, sim.NetID) {
	xorOut := b.AllocNet(1)
	andOut := b.AllocNet(1)

	xorNode := &XOR{InA: inA, InB: inB, Out: xorOut}
	b.AddNode(xorNode)

	andNode := &AND{InA: inA, InB: inB, Out: andOut}
	b.AddNode(andNode)

	return xorOut, andOut
}

func (b *Builder) FullAdder(inA, inB, inC sim.NetID) (sim.NetID, sim.NetID) {
	abAdderSum, abAdderCarry := b.HalfAdder(inA, inB)
	ccAdderSum, ccAdderCarry := b.HalfAdder(abAdderSum, inC)

	carryOr := b.OR(abAdderCarry, ccAdderCarry, 1)

	return ccAdderSum, carryOr
}
