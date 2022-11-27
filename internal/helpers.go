package internal

func isDivisibleBy(integer int, divisor int) bool {
	return integer%divisor == 0
}

func penniesToPounds(pennies int) float64 {
	return float64(pennies) / 100.0
}
