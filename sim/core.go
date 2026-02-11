package sim

/* TYPE DEFINITIONS */

type NetID int
type NodeID int

type Value uint32
type LogicState uint8

const (
	L0 LogicState = 0 // 00
	L1 LogicState = 1 // 01
	LX LogicState = 2 // 10
	LZ LogicState = 3 // 11
)

type Net struct {
	Width uint8
	Val   Value
}

type Node interface {
	Eval(nets []Net)
}

type NodeIO interface {
	Node
	Inputs() []NetID
	Outputs() []NetID
}

type NetList struct {
	Nets       []Net
	Nodes      []Node
	EvalOrder  []NodeID
	ProbeNames map[NetID]string
	Probes     []Probe
}

type Simulator struct {
	NL    *NetList
	Cycle uint64
}
