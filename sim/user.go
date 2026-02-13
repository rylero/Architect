package sim

/* Input */
type Input struct {
	Out   NetID
	Val   Value
	Width uint8
}

func (g *Input) Eval(nets []Net) {
	nets[g.Out].Val = g.Val
	nets[g.Out].Width = g.Width
}

func (g *Input) Inputs() []NetID    { return nil }
func (g *Input) Outputs() []NetID   { return []NetID{g.Out} }
func (g *Input) IsSequential() bool { return false }

/* Probe */
type Probe struct {
	Loc  NetID
	Name string
}
