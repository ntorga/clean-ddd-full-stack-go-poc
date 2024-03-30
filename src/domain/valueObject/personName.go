package valueObject

import (
	"errors"
	"regexp"
)

const personNameRegex string = `^\p{L}[\p{L}\'\ \-]{2,100}$`

type PersonName string

func NewPersonName(value string) (PersonName, error) {
	re := regexp.MustCompile(personNameRegex)
	isValid := re.MatchString(value)
	if !isValid {
		return "", errors.New("InvalidPersonName")
	}

	return PersonName(value), nil
}

func NewPersonNamePanic(value string) PersonName {
	name, err := NewPersonName(value)
	if err != nil {
		panic(err)
	}
	return name
}

func (name PersonName) String() string {
	return string(name)
}
