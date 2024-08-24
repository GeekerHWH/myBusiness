package calculator

func average(nums ...float32) (res float32) {
	for _, num := range nums {
		res += num
	}
	return res / float32(len(nums))
}

// DiscountRate return usedPrice divided by newPrice (rate of discount)
func DiscountRate(newPrice float32, usedPrice ...float32) float32 {
	return average(usedPrice...) / newPrice
}
