package sim

/* Logic State Helpers */

func GetBit(v Value, i int) LogicState {
	return LogicState((v >> uint(i*2)) & 3)
}

func SetBit(v *Value, i int, state LogicState) {
	mask := Value(3) << uint(i*2)
	*v = (*v & ^mask) | Value(state)<<uint(i*2)
}

func CreateValue(v string) Value {
	val := Value(0)
	for i, char := range v {
		if char == '1' {
			SetBit(&val, i, L1)
		} else {
			SetBit(&val, i, L0)
		}
	}
	return val
}
