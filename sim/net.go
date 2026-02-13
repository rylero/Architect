package sim

import (
	"fmt"
	"reflect"
	"strings"
)

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

func isLatchCycle(nl NetList, cycle []NodeID) bool {
	if len(cycle) > 6 || len(cycle) < 2 { // latch cycles are small
		return false
	}

	notCount, orCount := 0, 0
	for _, nid := range cycle {
		nodeType := reflect.TypeOf(nl.Nodes[nid])
		typeName := nodeType.String()
		switch {
		case strings.HasSuffix(typeName, "*library.NOT") ||
			strings.HasSuffix(typeName, "*library.NOR"):
			notCount++
		case strings.HasSuffix(typeName, "*library.OR"):
			orCount++
		}
	}

	return notCount >= 1 && orCount >= 1
}

func findCycleStart(path []NodeID, backEdge NodeID) int {
	for i, node := range path {
		if node == backEdge {
			return i
		}
	}
	return -1
}

func dfsDetectCycle(curr NodeID, graph map[NodeID][]NodeID,
	visited, recStack map[NodeID]bool, nl NetList) bool {
	visited[curr] = true
	recStack[curr] = true

	for _, next := range graph[curr] {
		if recStack[next] {
			// Cycle found! Extract cycle from recStack
			return true
		}
		if !visited[next] {
			if dfsDetectCycle(next, graph, visited, recStack, nl) {
				return true
			}
		}
	}

	recStack[curr] = false
	return false
}

func CreateEvalOrder(nl NetList) []NodeID {
	dependencies, indegrees := BuildDependencyGraph(nl)

	graph := make(map[NodeID][]NodeID, len(dependencies))
	for i, neighbors := range dependencies {
		graph[NodeID(i)] = neighbors
	}

	var combNodes, seqNodes []NodeID
	for i := range nl.Nodes {
		if nl.Nodes[i].(Node).IsSequential() {
			seqNodes = append(seqNodes, NodeID(i))
		} else {
			combNodes = append(combNodes, NodeID(i))
		}
	}

	visited := make(map[NodeID]bool)
	recStack := make(map[NodeID]bool)

	for _, nid := range combNodes {
		if !visited[nid] {
			if dfsDetectCycle(nid, graph, visited, recStack, nl) {
				fmt.Println("Latch cycle detected - allowing sequential logic")
			}
		}
	}

	// Kahn's algorithm on combinational nodes only
	nl.CombOrder = kahnSortComb(dependencies, indegrees, combNodes)
	nl.SeqOrder = seqNodes

	// Backward compatible full order
	nl.EvalOrder = append(nl.CombOrder, nl.SeqOrder...)
	return nl.EvalOrder
}

func kahnSortComb(succs [][]NodeID, indegrees []int, combNodes []NodeID) []NodeID {
	nComb := len(combNodes)

	// Track combinational nodes
	combSet := make(map[NodeID]bool)
	for _, nid := range combNodes {
		combSet[nid] = true
	}

	queue := []NodeID{}
	for _, nid := range combNodes {
		if indegrees[nid] == 0 {
			queue = append(queue, nid)
		}
	}

	order := []NodeID{}
	processed := 0

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		order = append(order, curr)
		processed++

		// Only combinational successors
		for _, next := range succs[curr] {
			if combSet[next] {
				indegrees[next]--
				if indegrees[next] == 0 {
					queue = append(queue, next)
				}
			}
		}
	}

	// Detect combinational cycles
	if processed != nComb {
		panic(fmt.Sprintf("Combinational cycle! Processed %d/%d comb nodes", processed, nComb))
	}

	return order
}
