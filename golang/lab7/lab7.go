package lab7

import "fmt"

func RunLab7() {
	cucumber := &Food{120, "огурцы", "пятерочка", 40}
	mascara := &Cosmetics{"тушь", 500, "Beauty Bomb"}
	jeans := &Clothes{"джинсы", 4999, "Colin`s", "хлопок", "все"}
	products := []Product{cucumber, mascara, jeans}
	fmt.Println("Список продуктов")
	fmt.Println("Стоимость покупок:", CalculateProductsSum(products), "рублей")
	for _, product := range products {
		fmt.Println(product.GetInfo())
	}

	for _, product := range products {
		product.setdiscount(40)
	}
	fmt.Println("Стоимость покупок после применения скидки 40%:", CalculateProductsSum(products), "рублей")

	fmt.Println("Информация по товару джинсы")
	fmt.Println("материал джинс:", jeans.getMaterial())
	fmt.Println("сезон, во время которого их носят:", jeans.getSeason())
	fmt.Println("бренд джинс:", jeans.getBrand())
}
