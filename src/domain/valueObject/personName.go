package valueObject

import (
	"errors"
	"regexp"
	"strings"

	voHelper "github.com/ntorga/clean-ddd-full-stack-go-poc/src/domain/valueObject/helper"
)

const personNameRegex string = `^\p{L}[\p{L}\'\ \-]{2,100}$`

type PersonName string

func NewPersonName(value interface{}) (PersonName, error) {
	stringValue, err := voHelper.InterfaceToString(value)
	if err != nil {
		return "", errors.New("PersonNameMustBeString")
	}

	stringValue = strings.TrimSpace(stringValue)

	re := regexp.MustCompile(personNameRegex)
	isValid := re.MatchString(stringValue)
	if !isValid {
		return "", errors.New("InvalidPersonName")
	}

	return PersonName(stringValue), nil
}

func (name PersonName) String() string {
	return string(name)
}
