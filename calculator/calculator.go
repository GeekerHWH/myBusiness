package calculator

func Average(nums ...int) (res float32) {
	for _, num := range nums {
		res += float32(num)
	}
	return res / float32(len(nums))
}

// DiscountRate return usedPrice divided by newPrice (rate of discount)
func DiscountRate(newPrice, usedPrice float32) float32 {
	return usedPrice / newPrice
}
