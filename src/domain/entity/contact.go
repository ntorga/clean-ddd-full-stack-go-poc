package entity

import "github.com/ntorga/clean-ddd-full-stack-go-poc/src/domain/valueObject"

type Contact struct {
	Id       valueObject.ContactId   `json:"id"`
	Name     valueObject.PersonName  `json:"name"`
	Nickname valueObject.Nickname    `json:"nickname"`
	Phone    valueObject.PhoneNumber `json:"phone"`
}

func NewContact(
	id valueObject.ContactId,
	name valueObject.PersonName,
	nickname valueObject.Nickname,
	phone valueObject.PhoneNumber,
) Contact {
	return Contact{
		Id:       id,
		Name:     name,
		Nickname: nickname,
		Phone:    phone,
	}
}

func (Contact) InitialEntities() []Contact {
	contactId, _ := valueObject.NewContactId(1)
	contactName, _ := valueObject.NewPersonName("Christopher Pike")
	contactNickname, _ := valueObject.NewNickname("Chris")
	contactPhone, _ := valueObject.NewPhoneNumber("555-17010")
	captainEntity := NewContact(contactId, contactName, contactNickname, contactPhone)

	contactId, _ = valueObject.NewContactId(2)
	contactName, _ = valueObject.NewPersonName("Una Chin-Riley")
	contactNickname, _ = valueObject.NewNickname("Una")
	contactPhone, _ = valueObject.NewPhoneNumber("555-17011")
	firstOfficerEntity := NewContact(contactId, contactName, contactNickname, contactPhone)

	contactId, _ = valueObject.NewContactId(3)
	contactName, _ = valueObject.NewPersonName("S'Chn T'Gai Spock")
	contactNickname, _ = valueObject.NewNickname("Spock")
	contactPhone, _ = valueObject.NewPhoneNumber("555-17012")
	scienceOfficerEntity := NewContact(contactId, contactName, contactNickname, contactPhone)

	return []Contact{
		captainEntity,
		firstOfficerEntity,
		scienceOfficerEntity,
	}
}
