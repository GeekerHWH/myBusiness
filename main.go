package main

import (
	"fmt"
	"myBusiness/calculator"
	"myBusiness/database"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	var name string
	var newPrice int
	nums := make([]int, 5)

	fmt.Scan(&name)
	fmt.Scan(&newPrice)
	for i := 0; i < len(nums); i++ {
		fmt.Scan(&nums[i])
	}
	//<--------------dealling input-------------->

	// calculate
	avg := calculator.Average(nums...)
	ratio := calculator.DiscountRate(float32(newPrice), avg)

	// 调试作用
	fmt.Println()
	fmt.Printf("二手均值\t折价率\n")
	fmt.Printf("%.4f\t%.4f\n", avg, ratio)

	//<--------------数据库功能开发ing...-------------->
	db, err := gorm.Open(sqlite.Open("myBusiness.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema(Create Table If Not Exist)
	err = db.AutoMigrate(&database.Goods{})
	if err != nil {
		panic("failed to migrate schema")
	}

	// Create a new Goods record
	good := database.Goods{
		Name:         name,
		NewPrice:     float32(newPrice),
		DiscountRate: ratio,
	}

	// Insert the record into the database
	err = db.Create(&good).Error
	if err != nil {
		panic("failed to insert record")
	}
	fmt.Println("Record inserted successfully")
}
