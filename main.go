package main

import (
	"fmt"
	"myBusiness/calculator"
)

func main() {
	var newPrice int
	fmt.Scan(&newPrice)

	nums := make([]int, 5)
	for i := 0; i < len(nums); i++ {
		fmt.Scan(&nums[i])
	}

	avg := calculator.Average(nums...)
	ratio := calculator.DiscountRate(float32(newPrice), avg)

	fmt.Println()
	fmt.Printf("二手均值\t折价率\n")
	fmt.Printf("%.4f\t%.4f\n", avg, ratio)
}
