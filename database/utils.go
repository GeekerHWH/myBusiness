package database

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
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

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "商品名称", "一手现货", "二手均值", "折价率"})
	for _, good := range goods {
		t.AppendRows([]table.Row{{good.ID, good.Name, good.NewPrice, good.NewPrice * good.DiscountRate, good.DiscountRate}})
	}
	t.AppendSeparator()
	t.AppendFooter(table.Row{"", "", "End", ""})
	t.Render()
}
