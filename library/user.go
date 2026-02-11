package library

import "architect/sim"

/* User Builder Functions */
func (b *Builder) Input(val sim.Value, width uint8, name string) sim.NetID {
	inputOut := b.AllocNamedNet(width, name)

	input := &sim.Input{Out: inputOut, Val: val, Width: width}
	b.AddNode(input)

	return inputOut
}
