package controller

import (
	"fmt"
	"mini_challenge_5_gorm/database"
	"mini_challenge_5_gorm/models"
)

func CreateVariant(name string, quantity int, productId int) {
	DB := database.GetDB()

	variant := models.Variant{
		VariantName:  name,
		Quantity: quantity,
		ProductID:  productId,
	}
	err := DB.Create(&variant).Error
	if err != nil {
		fmt.Println("can`t creating variant :", err)
		return
	}
	fmt.Printf("succes create variant %+v", variant)
}

func GetProductWithVariant() {
	DB := database.GetDB()
	products := models.Product{}
	err := DB.Preload("Variants").Find(&products).Error
	if err != nil {
		fmt.Println("can`t getting variants", err)
		return
	}
	fmt.Printf("product datas : %+v", products)
}

func  DeleteVariantByID(id int) {
	DB := database.DB
	variant := models.Variant{}
	err := DB.Where("id = ?", id).Delete(&variant).Error
	if err != nil {
		fmt.Println("can`t delete variant", err)
		return
	}

	fmt.Printf("succes delete variant with id %d", id)

}
