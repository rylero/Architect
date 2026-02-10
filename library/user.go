package library

import "architect/sim"

/* User Builder Functions */
func (b *Builder) Input(val sim.Value, width uint8) sim.NetID {
	inputOut := b.AllocNet(width)

	input := &sim.Input{Out: inputOut, Val: val}
	b.AddNode(input)

	return inputOut
}
