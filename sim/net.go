package sim

/* Node Connections */

type NodeConnections struct {
	Inputs  []NetID
	Outputs []NetID
}

func GetNodeConnections(nl NetList) []NodeConnections {
	conns := make([]NodeConnections, len(nl.Nodes))

	for i, node := range nl.Nodes {
		if nio, ok := node.(NodeIO); ok {
			conns[i].Inputs = nio.Inputs()
			conns[i].Outputs = nio.Outputs()
		}
	}
	return conns
}

func BuildDependencyGraph(nl NetList) ([][]NodeID, []int) {
	conns := GetNodeConnections(nl)
	n := len(nl.Nodes)

	netDrivers := make([][]NodeID, len(nl.Nets))
	netConsumers := make([][]NodeID, len(nl.Nets))

	for i := range conns {
		for _, out := range conns[i].Outputs {
			netDrivers[out] = append(netDrivers[out], NodeID(i))
		}
		for _, in := range conns[i].Inputs {
			netConsumers[in] = append(netConsumers[in], NodeID(i))
		}
	}

	succs := make([][]NodeID, n)
	indegrees := make([]int, n)

	for netID := range netDrivers {
		for _, driver := range netDrivers[netID] {
			for _, consumer := range netConsumers[netID] {
				if driver != consumer {
					succs[driver] = append(succs[driver], consumer)
					indegrees[consumer]++
				}
			}
		}
	}

	return succs, indegrees
}

func CreateEvalOrder(nl NetList) []NodeID {
	dependencies, indegrees := BuildDependencyGraph(nl)

	order := make([]NodeID, 0, len(nl.Nodes))
	readyQ := make([]NodeID, 0)

	for i := 0; i < len(indegrees); i++ {
		if indegrees[i] == 0 {
			readyQ = append(readyQ, NodeID(i))
		}
	}

	for len(readyQ) > 0 {
		u := readyQ[0]
		readyQ = readyQ[1:]

		order = append(order, u)

		for _, v := range dependencies[u] {
			indegrees[v]--
			if indegrees[v] == 0 {
				readyQ = append(readyQ, v)
			}
		}
	}

	if len(order) != len(nl.Nodes) {
		panic("Cycle detected in circut")
	}

	return order
}
