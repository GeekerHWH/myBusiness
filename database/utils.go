package database

import (
	"fmt"

	"gorm.io/gorm"
)

func createTable() {
	// Migrate the schema
	// err = db.AutoMigrate(&database.Goods{})
	// if err != nil {
	// 	panic("failed to migrate schema")
	// }
	// fmt.Println("Table created successfully")
}

func selectGoods() {
	// Select the record from the database
	// var selectedGood database.Goods
	// err = db.First(&selectedGood, 1).Error
	// if err != nil {
	// 	panic("failed to select record")
	// }
	// 调试作用
	// fmt.Println()
	// fmt.Printf("ID\t商品名称\t二手均值\t折价率\n")
	// fmt.Printf("%d\t%s\t%.4f\t%.4f\n", selectedGood.ID, selectedGood.Name, selectedGood.NewPrice, selectedGood.DiscountRate)
}

func InsertGood(db *gorm.DB, name string, newPrice float32, discountRate float32) {
	// Create a new Goods record
	good := Goods{
		Name:         name,
		NewPrice:     float32(newPrice),
		DiscountRate: discountRate,
	}

	// Insert the record into the database
	err := db.Create(&good).Error
	if err != nil {
		panic("failed to insert record")
	}
	fmt.Println("Record inserted successfully")
}

func PrintAll(db *gorm.DB) {
	// Select all records from the database
	var goods []Goods
	err := db.Find(&goods).Error
	if err != nil {
		panic("failed to select records")
	}

	// Print all records
	fmt.Println()
	fmt.Printf("商品名称\t二手均值\t折价率\n")
	for _, good := range goods {
		fmt.Printf("%s\t%.4f\t%.4f\n", good.Name, good.NewPrice, good.DiscountRate)
	}
}
