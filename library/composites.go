package library

import "architect/sim"

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
