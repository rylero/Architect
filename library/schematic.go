package library

import (
	"architect/sim"
	"fmt"
	"strings"
)

type SchematicOptions struct {
	ShowBusWidth bool
}

func ToDOT(nl *sim.NetList, opts SchematicOptions) string {
	var b strings.Builder
	b.WriteString("digraph circuit {\n")
	b.WriteString("  rankdir=LR;\n") // left→right

	// 1) Nodes: one per gate/module
	for i, node := range nl.Nodes {
		label := fmt.Sprintf("n%d", i)
		shape := "box"

		switch n := node.(type) {
		case *AND:
			label = "AND"
		case *OR:
			label = "OR"
		case *XOR:
			label = "XOR"
		case *Split:
			label = "SPLIT"
		case *Join:
			label = "JOIN"
		case *sim.Input:
			label = "IN"
			shape = "oval"
			// add more as needed
			_ = n
		}

		fmt.Fprintf(&b, "  node%d [label=\"%s\", shape=%s];\n", i, label, shape)
	}

	// 2) Edges: from driver node → sink node for each net
	// You already have NodeIO or similar for Inputs/Outputs. [cite:15]
	conns := sim.GetNodeConnections(*nl) // Inputs/Outputs per node

	// net → list of drivers and sinks
	type ds struct{ drivers, sinks []sim.NodeID }
	netMap := map[sim.NetID]*ds{}

	for nid, c := range conns {
		for _, n := range c.Outputs {
			if netMap[n] == nil {
				netMap[n] = &ds{}
			}
			netMap[n].drivers = append(netMap[n].drivers, sim.NodeID(nid))
		}
		for _, n := range c.Inputs {
			if netMap[n] == nil {
				netMap[n] = &ds{}
			}
			netMap[n].sinks = append(netMap[n].sinks, sim.NodeID(nid))
		}
	}

	for netID, dss := range netMap {
		for _, d := range dss.drivers {
			for _, s := range dss.sinks {
				if d == s {
					continue
				}
				label := ""
				if opts.ShowBusWidth {
					width := nl.Nets[netID].Width
					label = fmt.Sprintf(" [label=\"%d\"]", width)
				}
				fmt.Fprintf(&b, "  node%d -> node%d%s;\n", d, s, label)
			}
		}
	}

	b.WriteString("}\n")
	return b.String()
}
