package cmd

import (
	"fmt"
	"myBusiness/calculator"
	"myBusiness/database"

	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init() {
	rootCmd.AddCommand(insertGoods)
}

var insertGoods = &cobra.Command{
	Use:   "insert",
	Short: "Insert a new good information to database",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		var name string
		var newPrice int
		nums := make([]float32, 5)

		fmt.Println("商品名称 新货现价 5*二手价格")
		fmt.Scan(&name)
		fmt.Scan(&newPrice)
		for i := 0; i < len(nums); i++ {
			fmt.Scan(&nums[i])
		}
		// <--------------dealling input-------------->

		// calculate
		ratio := calculator.DiscountRate(float32(newPrice), nums...)

		// <--------------数据库功能开发ing...-------------->
		db, err := gorm.Open(sqlite.Open("myBusiness.db"), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}

		// Create Table If Not Exist
		err = db.AutoMigrate(&database.Goods{})
		if err != nil {
			panic("failed to migrate schema")
		}

		database.InsertGood(db, name, float32(newPrice), ratio)

		database.PrintAll(db)
	},
}
