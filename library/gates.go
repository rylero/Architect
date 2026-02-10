package library

import . "architect/sim"

/* AND Gate */

type AND struct {
	InA, InB NetID
	Out      NetID
}

func and(a, b LogicState) LogicState {
	if a == L0 || b == L0 {
		return L0
	} // Dominant 0
	if a == L1 && b == L1 {
		return L1
	} // Both true
	if a == LX || b == LX {
		return LX
	} // Contamination
	return LZ // Z/no driver
}

func (g *AND) Eval(nets []Net) {
	width := nets[g.InA].Width

	for i := 0; i < int(width); i++ {
		a := GetBit(nets[g.InA].Val, i)
		b := GetBit(nets[g.InB].Val, i)

		result := and(a, b)
		SetBit(&nets[g.Out].Val, i, result)
	}
}

func (g *AND) Inputs() []NetID  { return []NetID{g.InA, g.InB} }
func (g *AND) Outputs() []NetID { return []NetID{g.Out} }

/* Not */

type NOT struct {
	In  NetID
	Out NetID
}

func not(v LogicState) LogicState {
	if v == L1 {
		return L0
	}
	if v == L0 {
		return L1
	}
	if v == LZ {
		return LZ
	}
	return LX
}

func (g *NOT) Eval(nets []Net) {
	width := nets[g.In].Width

	for i := 0; i < int(width); i++ {
		v := GetBit(nets[g.In].Val, i)

		result := not(v)
		SetBit(&nets[g.Out].Val, i, result)
	}
}

func (g *NOT) Inputs() []NetID  { return []NetID{g.In} }
func (g *NOT) Outputs() []NetID { return []NetID{g.Out} }

/* OR Gate */

type OR struct {
	InA, InB NetID
	Out      NetID
}

func or(a, b LogicState) LogicState {
	if a == L1 || b == L1 {
		return L1
	} // Dominant 1
	if a == L0 && b == L0 {
		return L0
	} // Both false
	if a == LX || b == LX {
		return LX
	}
	return LZ
}

func (g *OR) Eval(nets []Net) {
	width := nets[g.InA].Width

	for i := 0; i < int(width); i++ {
		a := GetBit(nets[g.InA].Val, i)
		b := GetBit(nets[g.InB].Val, i)

		result := or(a, b)
		SetBit(&nets[g.Out].Val, i, result)
	}
}

func (g *OR) Inputs() []NetID  { return []NetID{g.InA, g.InB} }
func (g *OR) Outputs() []NetID { return []NetID{g.Out} }

/* XOR Gate */
type XOR struct {
	InA, InB NetID
	Out      NetID
}

func xor(a, b LogicState) LogicState {
	if (a == L0 && b == L0) || (a == L1 && b == L1) {
		return L0
	} // Both false or both false
	if a == L1 || b == L1 {
		return L1
	}
	if a == LX || b == LX {
		return LX
	}
	return LZ
}

func (g *XOR) Eval(nets []Net) {
	width := nets[g.InA].Width

	for i := 0; i < int(width); i++ {
		a := GetBit(nets[g.InA].Val, i)
		b := GetBit(nets[g.InB].Val, i)

		result := xor(a, b)
		SetBit(&nets[g.Out].Val, i, result)
	}
}

func (g *XOR) Inputs() []NetID  { return []NetID{g.InA, g.InB} }
func (g *XOR) Outputs() []NetID { return []NetID{g.Out} }
