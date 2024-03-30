package dbModel

import (
	"time"

	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/domain/entity"
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/domain/valueObject"
)

type Contact struct {
	ID        uint   `gorm:"primarykey"`
	Name      string `gorm:"not null"`
	Nickname  string `gorm:"not null"`
	Phone     string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Contact) TableName() string {
	return "contacts"
}

func (model Contact) InitialEntries() []interface{} {
	initialEntities := entity.Contact{}.InitialEntities()

	initialEntries := []interface{}{}
	for _, entity := range initialEntities {
		entry, _ := model.ToModel(entity)
		initialEntries = append(initialEntries, entry)
	}

	return initialEntries
}

func (Contact) ToModel(entity entity.Contact) (Contact, error) {
	return Contact{
		ID:       uint(entity.Id),
		Name:     entity.Name.String(),
		Nickname: entity.Nickname.String(),
		Phone:    entity.Phone.String(),
	}, nil
}

func (model Contact) ToEntity() (entity.Contact, error) {
	var contactEntity entity.Contact

	contactId, err := valueObject.NewContactId(model.ID)
	if err != nil {
		return contactEntity, err
	}

	name, err := valueObject.NewPersonName(model.Name)
	if err != nil {
		return contactEntity, err
	}

	nickname, err := valueObject.NewNickname(model.Nickname)
	if err != nil {
		return contactEntity, err
	}

	phone, err := valueObject.NewPhoneNumber(model.Phone)
	if err != nil {
		return contactEntity, err
	}

	return entity.NewContact(
		contactId,
		name,
		nickname,
		phone,
	), nil
}
