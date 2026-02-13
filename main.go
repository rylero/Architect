package main

import (
	"architect/library"
	"architect/sim"
	"fmt"
)

func main() {
	b := library.NewBuilder()
	b.EnterScope("top")
	d := b.Input(sim.CreateValue("1"), 1, "D")
	en := b.Input(sim.CreateValue("1"), 1, "EN")

	b.DLatch(d, en)

	nl := b.Build()
	sim := sim.Simulator{NL: nl}
	sim.Step()

	probes := sim.ReadProbes()
	fmt.Println("Q:", probes["Q"], "Qbar:", probes["Qbar"])
}
