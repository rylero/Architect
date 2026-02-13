package sim

import "fmt"

func (s *Simulator) Step() {
	// Phase 1: Combinational stabilization (iterative eval)
	maxIters := 32
	for iter := 0; iter < maxIters; iter++ {
		changed := false

		for _, nid := range s.NL.CombOrder {
			oldVal := s.NL.Nets[nid].Val

			s.NL.Nodes[nid].Eval(s.NL.Nets)

			if oldVal != s.NL.Nets[nid].Val {
				changed = true
			}
		}

		if !changed {
			break
		}
	}

	// Phase 2: Sequential once (latches settle in next comb iter)
	for _, nid := range s.NL.SeqOrder {
		s.NL.Nodes[nid].Eval(s.NL.Nets)
	}
}

func FormatValue(val Value, width uint8) string {
	stateStr := make([]byte, width)
	for i := uint8(0); i < width; i++ {
		state := GetBit(val, int(i))
		switch state {
		case L0:
			stateStr[i] = '0'
		case L1:
			stateStr[i] = '1'
		case LX:
			stateStr[i] = 'X'
		case LZ:
			stateStr[i] = 'Z'
		default:
			stateStr[i] = '?'
		}
	}
	return string(stateStr)
}

func (s *Simulator) ReadProbes() map[string]string {
	result := make(map[string]string)
	for _, p := range s.NL.Probes {
		val := s.NL.Nets[p.Loc].Val
		name := s.NL.ProbeNames[p.Loc]
		if name == "" {
			name = fmt.Sprintf("net%d", p.Loc)
		}
		result[name] = FormatValue(val, s.NL.Nets[p.Loc].Width)
	}
	return result
}
