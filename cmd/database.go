package cmd

import (
	"fmt"
	"myBusiness/calculator"
	"myBusiness/database"
	"os"

	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var name string
var newPrice float32

func init() {
	rootCmd.AddCommand(insertGoods)
	rootCmd.AddCommand(deleteDB)

	insertGoods.Flags().StringVarP(&name, "name", "n", "", "Goods' name")
	insertGoods.Flags().Float32VarP(&newPrice, "price", "p", 0, "new price of the good")
}

var insertGoods = &cobra.Command{
	Use:   "insert",
	Short: "Insert a new good information to database",
	Long:  `Insert and calculate new goods' information to database`,
	Run: func(cmd *cobra.Command, args []string) {

		nums := make([]float32, 5)

		fmt.Println("请输入5个随机二手价格样本（空格分割）")
		for i := 0; i < len(nums); i++ {
			fmt.Scan(&nums[i])
		}
		// <--------------dealling input-------------->

		// calculate
		ratio := calculator.DiscountRate(newPrice, nums...)

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

		database.InsertGood(db, name, newPrice, ratio)

		database.PrintAll(db)
	},
}

var deleteDB = &cobra.Command{
	Use:   "delete",
	Short: "Delete database",
	Long:  `Delete myBusiness.db`,
	Run: func(cmd *cobra.Command, args []string) {
		err := os.Remove("./myBusiness.db")
		if err != nil {
			panic("Failed to delete myBusiness.db\n")
		}
	},
}
