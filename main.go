package main

import (
	"architect/library"
	"architect/sim"
	"fmt"
)

func main() {
	b := library.NewBuilder()

	inputA := b.Input(sim.CreateValue("1"), 1)
	inputB := b.Input(sim.CreateValue("1"), 1)
	inputC := b.Input(sim.CreateValue("1"), 1)

	sum, carry := b.FullAdder(inputA, inputB, inputC)

	nl := b.Build()

	simulator := sim.Simulator{NL: nl}
	simulator.Probes = []sim.Probe{{Loc: sum, Name: "sum.out"}, {Loc: carry, Name: "carry.out"}}

	simulator.Step()
	probes := simulator.ReadProbes()

	fmt.Println("sum", probes["sum.out"])
	fmt.Println("carry", probes["carry.out"])
}
