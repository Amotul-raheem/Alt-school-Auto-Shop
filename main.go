package main

import (
	"fmt"
	"log"
)

type Car struct {
	model     string
	brandName string
}

type Product struct {
	Car
	quantity int
	price    int
	id       int
}

type Store struct {
	allProducts       []Product
	soldProducts      []Product
	availableProducts []Product
}

func (p *Product) isInStock() bool {
	return p.quantity != 0
}

func (s *Store) addProductToStore(p Product) {
	s.allProducts = append(s.allProducts, p)
	s.availableProducts = append(s.availableProducts, p)
	log.Println("The Item below was added to store")
	s.displayProducts([]Product{p})

}

func (s *Store) sellProductInStore(id int) {
	for index, product := range s.allProducts {
		if product.id == id {
			s.soldProducts = append(s.soldProducts, product)
			s.availableProducts = append(s.availableProducts[:index], s.availableProducts[index+1:]...)
			log.Println("The Item below has been sold")
			s.displayProducts([]Product{product})
		}
	}
}

func (s *Store) getTotalPriceOfItemsSold() int {
	totalSales := 0
	for _, product := range s.soldProducts {
		totalSales += product.price
	}
	return totalSales
}

func (s *Store) displayProductsSold() {
	s.displayProducts(s.soldProducts)
	log.Printf("Total Sales Price : %d", s.getTotalPriceOfItemsSold())
}

func (s *Store) displayProductsInStore() {
	s.displayProducts(s.allProducts)
}

func (s Store) getNumberOfProductsInStore() int {
	return len(s.availableProducts)
}

func (s Store) displayProducts(products []Product) {
	for _, product := range products {
		log.Printf("Product Id : %d", product.id)
		log.Printf("Product Price : %d", product.price)
		log.Printf("Product Name : %s", product.Car.brandName)
		log.Printf("Product Model : %s \n", product.Car.model)
		fmt.Println("\n")
	}
}

func main() {
	//Create Cars and Add to Products

	//Tesla
	teslaModel3Car := Car{
		model:     "Model 3",
		brandName: "Tesla",
	}
	productTeslaModel3 := Product{
		Car:      teslaModel3Car,
		quantity: 5,
		price:    40000,
		id:       1,
	}
	teslaModelYCar := Car{
		model:     "Model Y",
		brandName: "Tesla",
	}
	productTeslaModelY := Product{
		Car:      teslaModelYCar,
		quantity: 2,
		price:    20000,
		id:       2,
	}
	//Toyota
	toyotaCorolla := Car{
		model:     "Corolla",
		brandName: "Toyota",
	}
	productToyotaCorolla := Product{
		Car:      toyotaCorolla,
		quantity: 8,
		price:    10000,
		id:       3,
	}
	toyotaCamry := Car{
		model:     "Camry",
		brandName: "Toyota",
	}
	productToyotaCamry := Product{
		Car:      toyotaCamry,
		quantity: 8,
		price:    10000,
		id:       4,
	}

	// Create Store and add Products
	store := new(Store)
	//Add Tesla products
	store.addProductToStore(productTeslaModelY)
	store.addProductToStore(productTeslaModel3)
	//Add Toyota products
	store.addProductToStore(productToyotaCorolla)

	//1) Get number of items in store
	log.Printf("There are %d number of products in store \n", store.getNumberOfProductsInStore())
	store.addProductToStore(productToyotaCamry)
	log.Printf("There are %d number of products in store \n", store.getNumberOfProductsInStore())

	//2) List all product items in store
	store.displayProductsInStore()

	//3) Sell products by productId
	store.sellProductInStore(3)
	store.sellProductInStore(1)
	log.Printf("There are %d number of products in store", store.getNumberOfProductsInStore())

	//4) Display list of sold items
	store.displayProductsSold()

}
