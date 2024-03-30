package repository

import "github.com/ntorga/clean-ddd-taghs-poc-contacts/src/domain/dto"

type ContactCmdRepo interface {
	Add(dto dto.AddContact) error
}
