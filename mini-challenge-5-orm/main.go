package main

import (
	"mini_challenge_5_gorm/controller"
	"mini_challenge_5_gorm/database"
)

func main() {
	database.ConnectDB()


	controller.CreateProduct("laptop")
	controller.UpdateProduct(2,"table")
	controller.GetProductByID(1)
	controller.DeleteProductByID(2)
	controller.CreateVariant("asus notebook 01",4,1)
	controller.GetProductWithVariant()
	controller.DeleteVariantByID(1)
}
