package sim

import "fmt"

func (s *Simulator) Step() {
	for _, id := range s.NL.EvalOrder {
		s.NL.Nodes[id].Eval(s.NL.Nets)
	}
	s.Cycle++
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
