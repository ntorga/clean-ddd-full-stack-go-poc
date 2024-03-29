package infra

import "github.com/ntorga/clean-ddd-taghs-poc-contacts/src/domain/entity"

type ContactQueryRepo struct {
}

func NewContactQueryRepo() *ContactQueryRepo {
	return &ContactQueryRepo{}
}

func (repo *ContactQueryRepo) Get() ([]entity.Contact, error) {
	return []entity.Contact{}, nil
}
