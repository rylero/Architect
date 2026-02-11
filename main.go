package main

import (
	"architect/library"
	"architect/sim"
	"fmt"
)

func main() {
	b := library.NewBuilder()

	inputA := b.Input(sim.CreateValue("10010010"), 8)
	inputB := b.Input(sim.CreateValue("11001001"), 8)
	inputC := b.Input(sim.CreateValue("0"), 1)

	sum, carry := b.RippleCarryAdder(inputA, inputB, inputC, 8)

	nl := b.Build()

	simulator := sim.Simulator{NL: nl}
	simulator.Probes = []sim.Probe{{Loc: sum, Name: "sum.out"}, {Loc: carry, Name: "carry.out"}}

	simulator.Step()
	probes := simulator.ReadProbes()

	fmt.Println("sum", probes["sum.out"])
	fmt.Println("carry", probes["carry.out"])

	fmt.Println(library.ToDOT(nl, library.SchematicOptions{ShowBusWidth: true}))
}
