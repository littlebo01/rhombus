package rhombus

func Chain(strategies ...Strategy) Strategy {
	return &batchStrategy{
		size: 1,
		strategies: strategies,
	}
}
