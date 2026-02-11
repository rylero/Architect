package main

import (
	"architect/library"
	"architect/sim"
	"fmt"
)

func main() {
	b := library.NewBuilder()
	b.EnterScope("top")

	inputA := b.Input(sim.CreateValue("10010010"), 8, "InputA")
	inputB := b.Input(sim.CreateValue("11001001"), 8, "InputB")
	inputC := b.Input(sim.CreateValue("0"), 1, "InputC")

	b.RippleCarryAdder(inputA, inputB, inputC, 8)

	nl := b.Build()

	simulator := sim.Simulator{NL: nl}

	simulator.Step()
	probes := simulator.ReadProbes()

	fmt.Println("A+B=3+5:", probes["top.InputA"], "+", probes["top.InputB"], "=", probes["top.ripple.join.out"])
}
