package sim

func (s *Simulator) Step() {
	for _, id := range s.NL.EvalOrder {
		s.NL.Nodes[id].Eval(s.NL.Nets)
	}
	s.Cycle++
}

func (sim *Simulator) ReadProbes() map[string]string {
	result := make(map[string]string)
	for _, p := range sim.Probes {
		net := sim.NL.Nets[p.Loc]
		val := net.Val
		width := net.Width
		stateStr := make([]byte, width)

		for i := uint8(0); i < width; i++ {
			shift := i * 2
			state := LogicState((val >> shift) & 3)
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
		result[p.Name] = string(stateStr)
	}
	return result
}
