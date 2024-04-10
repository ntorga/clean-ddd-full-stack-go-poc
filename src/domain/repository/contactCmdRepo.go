package repository

import (
	"github.com/ntorga/clean-ddd-full-stack-go-poc/src/domain/dto"
	"github.com/ntorga/clean-ddd-full-stack-go-poc/src/domain/valueObject"
)

type ContactCmdRepo interface {
	Create(dto dto.CreateContact) error
	Update(dto dto.UpdateContact) error
	Delete(id valueObject.ContactId) error
}
