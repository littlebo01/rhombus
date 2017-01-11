package rhombus

func Multi(strategies ...Strategy) Strategy {
	return &batchStrategy{
		size: len(strategies),
		strategies: strategies,
	}
}
