package infrastructure

import "github.com/inventoryServices/domain"

type Repository interface {
	GetByShopCategory(string, uint) (domain.Product, error)
	GetName(string, uint) ([]domain.Product, error)
	GetByID(string) (domain.Product, error)
	AddItem(...*domain.Product) error
	SetToCategory(...*domain.Product) error
	CreateCategory(string) error
	Len() (uint, error)
	Clear() error
	DeleteItem(domain.Product) error
}

type Implementation int

const (
	Json Implementation = iota
)
