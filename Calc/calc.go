package calc

type Calc struct {
	Data1 float64
	Data2 float64
	Opr   string
}

func (cl Calc) AddNumber() float64 {
	return cl.Data1 + cl.Data2
}

func (cl Calc) SubstractNumber() float64 {
	return cl.Data1 - cl.Data2
}

func (cl Calc) MultiplyNumber() float64 {
	return cl.Data1 * cl.Data2
}

func (cl Calc) DivideNumber() float64 {
	return cl.Data1 / cl.Data2
}
