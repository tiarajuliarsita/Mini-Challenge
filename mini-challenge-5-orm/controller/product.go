package controller

import (
	"errors"
	"fmt"
	"mini_challenge_5_gorm/database"
	"mini_challenge_5_gorm/models"

	"gorm.io/gorm"
)

func CreateProduct(name string) {
	db := database.GetDB()

	product := models.Product{
		Name: name,
	}
	err := db.Create(&product).Error
	if err != nil {
		panic(err)
	}

	fmt.Println("success create product")
	fmt.Printf("%+v\n", product)
}

func UpdateProduct(id int, name string) {
	db := database.GetDB()
	product := models.Product{
		Name: name,
	}

	err := db.Model(&models.Product{}).Where("id = ?", id).Updates(&product).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("product not found")
		}
		panic(err)
	}

	fmt.Printf("succes updated product  to %s", name)
}

func GetProductByID(id int) {
	db := database.GetDB()

	var product models.Product
	err:= db.Find(&product, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("product not found")
		}
		panic(err)
	}
	fmt.Printf("product : %+v\n", product)
}

func  DeleteProductByID(id int) {
	DB := database.DB
	product := models.Product{}
	err := DB.Where("id = ?", id).Delete(&product).Error
	if err != nil {
		fmt.Println("can`t delete product", err)
		return
	}

	fmt.Println("succes deleted product")

}
