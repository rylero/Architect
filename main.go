package main

import (
	"architect/library"
	"architect/sim"
	"fmt"
)

func main() {
	b := library.NewBuilder()

	inputA := b.Input(sim.CreateValue("10110010"), 8)
	inputB := b.Input(sim.CreateValue("10011001"), 8)

	nand := b.NAND(inputA, inputB, 8)

	nl := b.Build()

	simulator := sim.Simulator{NL: nl}
	simulator.Probes = []sim.Probe{{Loc: nand, Name: "nand.out"}}

	simulator.Step()
	probes := simulator.ReadProbes()

	fmt.Println(probes["nand.out"]) // ~AND result: e.g., "01011011"
}
