package rhombus

func Multi(strategies ...Strategy) Strategy {
	size := len(strategies)

	return &batchStrategy{
		size: len(strategies),
		strategies: strategies,
	}
}
