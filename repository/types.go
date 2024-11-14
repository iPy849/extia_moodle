package repository

type Repository[T any, I any] interface {
	GetById(id I) (*T, error)
	Create() error
	Update() error
	Delete() error
}
