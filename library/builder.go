package library

import "architect/sim"

/* Builder */

type Builder struct {
	nl       *sim.NetList
	nextNet  sim.NetID
	nextNode sim.NodeID
}

func NewBuilder() *Builder {
	return &Builder{
		nl: &sim.NetList{
			Nets:      make([]sim.Net, 0),
			Nodes:     make([]sim.Node, 0),
			EvalOrder: make([]sim.NodeID, 0),
		},
		nextNet:  0,
		nextNode: 0,
	}
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
