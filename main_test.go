package main

import (
	"architect/library"
	"architect/sim"
	"testing"
)

func RunRippleCarryAdder(num1, num2, carry string) map[string]string {
	b := library.NewBuilder()
	b.EnterScope("top")

	inputA := b.Input(sim.CreateValue(num1), 8, "InputA")
	inputB := b.Input(sim.CreateValue(num2), 8, "InputB")
	inputC := b.Input(sim.CreateValue(carry), 1, "InputC")

	b.RippleCarryAdder(inputA, inputB, inputC, 8)

	simulator := sim.Simulator{NL: b.Build()}

	simulator.Step()

	probes := simulator.ReadProbes()

	return probes
}

func TestRippleCarryAdd(t *testing.T) {
	probes := RunRippleCarryAdder("10010010", "11001001", "0")

	if result := probes["top.ripple.join.out"]; result != "00111011" {
		t.Errorf(`RippleCarryAdder  = %q, %v, want "", error`, result, "00111011")
	}
}

func TestRippleCarryAddCarry(t *testing.T) {
	probes := RunRippleCarryAdder("11111111", "10000000", "0")

	result := probes["top.ripple.join.out"]
	carry := probes["top.ripple.7.fullAdder.carry"]

	if result != "00000000" {
		t.Errorf(`RippleCarryAdder with Carry = %q (with carry %q), want %q (with carry %q), error`, result, carry, "00000000", "1")
	}
}

func TestRippleCarryAddCarryIn(t *testing.T) {
	probes := RunRippleCarryAdder("11111111", "00000000", "1")

	result := probes["top.ripple.join.out"]
	carry := probes["top.ripple.7.fullAdder.carry"]

	if result != "00000000" {
		t.Errorf(`RippleCarryAdder with Carry In = %q (with carry %q), want %q (with carry %q), error`, result, carry, "00000000", "1")
	}
}
