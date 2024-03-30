package repository

import "github.com/ntorga/clean-ddd-taghs-poc-contacts/src/domain/dto"

type ContactCmdRepo interface {
	Create(dto dto.CreateContact) error
	Update(dto dto.UpdateContact) error
}
