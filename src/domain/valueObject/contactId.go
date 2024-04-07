package valueObject

import (
	"errors"
	"strconv"

	voHelper "github.com/ntorga/clean-ddd-taghs-poc-contacts/src/domain/valueObject/helper"
)

type ContactId uint64

func NewContactId(value interface{}) (ContactId, error) {
	id, err := voHelper.InterfaceToUint(value)
	if err != nil {
		return 0, errors.New("InvalidContactId")
	}

	return ContactId(id), nil
}

func (id ContactId) Uint() uint64 {
	return uint64(id)
}

func (id ContactId) String() string {
	return strconv.FormatUint(uint64(id), 10)
}
