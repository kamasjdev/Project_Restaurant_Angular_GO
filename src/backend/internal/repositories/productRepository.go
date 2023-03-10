package repositories

import (
	"fmt"

	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities"
)

type ProductRepository interface {
	Add(*entities.Product) error
	Update(entities.Product) error
	Delete(entities.Product) error
	Get(int64) (*entities.Product, error)
	GetAll() ([]entities.Product, error)
}

type inMemoryProductRepository struct {
	products []entities.Product
}

func NewInMemoryProductRepository() ProductRepository {
	return &inMemoryProductRepository{
		products: make([]entities.Product, 0),
	}
}

func (repo *inMemoryProductRepository) Add(product *entities.Product) error {
	var length int = len(repo.products)
	if length == 0 {
		product.Id = 1
		repo.products = append(repo.products, *product)
		return nil
	}

	lastElement := repo.products[length-1]
	product.Id = lastElement.Id + 1
	repo.products = append(repo.products, *product)
	return nil
}

func (repo *inMemoryProductRepository) Update(productToUpdate entities.Product) error {
	for index, product := range repo.products {
		if product.Id == productToUpdate.Id {
			product.Name = productToUpdate.Name
			product.Price = productToUpdate.Price
			product.Deleted = productToUpdate.Deleted
			product.Description = productToUpdate.Description
			repo.products[index] = product
		}
	}
	return fmt.Errorf("'Product' with id %v was not found", productToUpdate.Id)
}

func (repo *inMemoryProductRepository) Delete(productToDelete entities.Product) error {
	for index, product := range repo.products {
		if product.Id == productToDelete.Id {
			repo.products = append(repo.products[:index], repo.products[index+1:]...)
			return nil
		}
	}

	return fmt.Errorf("'Product' with id %v was not found", productToDelete.Id)
}

func (repo *inMemoryProductRepository) Get(id int64) (*entities.Product, error) {
	for _, product := range repo.products {
		if product.Id == id {
			return &product, nil
		}
	}

	return nil, fmt.Errorf("'Product' with id %v was not found", id)
}

func (repo *inMemoryProductRepository) GetAll() ([]entities.Product, error) {
	products := make([]entities.Product, 0)

	for _, product := range repo.products {
		if !product.Deleted {
			products = append(products, product)
		}
	}

	return products, nil
}
