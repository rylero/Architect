package sim

/* Input */
type Input struct {
	Out NetID
	Val Value
}

func (g *Input) Eval(nets []Net) {
	nets[g.Out].Val = g.Val
	nets[g.Out].Width = 16
}

func (g *Input) Inputs() []NetID  { return nil }
func (g *Input) Outputs() []NetID { return []NetID{g.Out} }

/* Probe */
type Probe struct {
	Loc  NetID
	Name string
}
