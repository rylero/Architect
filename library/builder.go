package library

import (
	"architect/sim"
	"fmt"
	"strings"
)

/* Builder */

type Builder struct {
	nl       *sim.NetList
	nextNet  sim.NetID
	nextNode sim.NodeID
	scopes   []string
}

func NewBuilder() *Builder {
	return &Builder{
		nl: &sim.NetList{
			Nets:       make([]sim.Net, 0),
			Nodes:      make([]sim.Node, 0),
			EvalOrder:  make([]sim.NodeID, 0),
			Probes:     make([]sim.Probe, 0),
			ProbeNames: make(map[sim.NetID]string),
		},
		nextNet:  0,
		nextNode: 0,
	}
}

func (b *Builder) EnterScope(name string) {
	b.scopes = append(b.scopes, name)
}

func (b *Builder) ExitScope() {
	b.scopes = b.scopes[:len(b.scopes)-1]
}

func (b *Builder) AllocNamedNet(width uint8, localName string) sim.NetID {
	id := b.AllocNet(width)
	fullName := strings.Join(b.scopes, ".") + "." + localName
	b.nl.ProbeNames[id] = fullName
	b.nl.Probes = append(b.nl.Probes, sim.Probe{Loc: id, Name: fullName})
	return id
}

func (b *Builder) AllocNet(width uint8) sim.NetID {
	id := b.nextNet
	b.nl.Nets = append(b.nl.Nets, sim.Net{Width: width})
	b.nextNet++
	return id
}

func (b *Builder) AddNode(node sim.Node) sim.NodeID {
	id := b.nextNode
	b.nl.Nodes = append(b.nl.Nodes, node)
	b.nextNode++
	return id
}

func (b *Builder) Build() *sim.NetList {
	b.nl.EvalOrder = sim.CreateEvalOrder(*b.nl)
	return b.nl
}

/* Split and Join */

type Split struct {
	In     sim.NetID   // Wide input net
	Outs   []sim.NetID // Output nets (must match NumOut, each width = In.Width / NumOut)
	Width  uint8       // Total bits (In.Width)
	NumOut uint8       // Number of output slices
}

func (s *Split) Eval(nets []sim.Net) {
	inVal := nets[s.In].Val
	outWidth := s.Width / s.NumOut
	for outIdx := uint8(0); outIdx < s.NumOut; outIdx++ {
		outNet := s.Outs[outIdx]
		outVal := sim.Value(0)
		for bit := uint8(0); bit < outWidth; bit++ {
			globalBit := outIdx*outWidth + bit
			sim.SetBit(&outVal, int(bit), sim.GetBit(inVal, int(globalBit)))
		}
		nets[outNet].Val = outVal
	}
}

type Join struct {
	Ins   []sim.NetID // Input nets (must match NumIn)
	Out   sim.NetID   // Wide output net
	Width uint8       // Total bits (Out.Width)
	NumIn uint8       // Number of input slices
}

func (j *Join) Eval(nets []sim.Net) {
	outVal := sim.Value(0)
	inWidth := j.Width / j.NumIn
	for inIdx := uint8(0); inIdx < j.NumIn; inIdx++ {
		inNet := j.Ins[inIdx]
		inVal := nets[inNet].Val
		for bit := uint8(0); bit < inWidth; bit++ {
			globalBit := inIdx*inWidth + bit
			sim.SetBit(&outVal, int(globalBit), sim.GetBit(inVal, int(bit)))
		}
	}
	nets[j.Out].Val = outVal
}

func (s *Split) Inputs() []sim.NetID  { return []sim.NetID{s.In} }
func (s *Split) Outputs() []sim.NetID { return s.Outs }

func (j *Join) Inputs() []sim.NetID  { return j.Ins }
func (j *Join) Outputs() []sim.NetID { return []sim.NetID{j.Out} }

// Scoped Split: wide bus → named slices
func (b *Builder) Split(in sim.NetID, numOut uint8) []sim.NetID {
	width := b.nl.Nets[in].Width
	if width%numOut != 0 {
		panic(fmt.Sprintf("Split: width %d %% %d != 0", width, numOut))
	}
	sliceWidth := width / numOut

	b.EnterScope("split")
	defer b.ExitScope()

	outs := make([]sim.NetID, numOut)
	for i := uint8(0); i < numOut; i++ {
		outs[i] = b.AllocNamedNet(sliceWidth, fmt.Sprintf("out%d", i))
	}

	node := &Split{In: in, Outs: outs, Width: width, NumOut: numOut}
	b.AddNode(node)
	return outs
}

// Scoped Join: slices → wide bus
func (b *Builder) Join(ins []sim.NetID, width uint8) sim.NetID {
	b.EnterScope("join")
	defer b.ExitScope()

	out := b.AllocNamedNet(width, "out")
	numIn := uint8(len(ins))

	node := &Join{Ins: ins, Out: out, Width: width, NumIn: numIn}
	b.AddNode(node)
	return out
}

func (g *AND) IsSequential() bool   { return false }
func (g *OR) IsSequential() bool    { return false }
func (g *NOT) IsSequential() bool   { return false }
func (g *XOR) IsSequential() bool   { return false }
func (g *WIRE) IsSequential() bool  { return false }
func (s *Split) IsSequential() bool { return false }
func (j *Join) IsSequential() bool  { return false }
