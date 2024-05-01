package valueObject

import (
	"errors"
	"regexp"
	"strings"

	voHelper "github.com/ntorga/clean-ddd-full-stack-go-poc/src/domain/valueObject/helper"
)

const phoneNumberRegex string = `^\(?\d{1,9}\)?[ \-\.]?\d{1,9}[ \-\.]?\d{1,9}$`

type PhoneNumber string

func NewPhoneNumber(value interface{}) (PhoneNumber, error) {
	stringValue, err := voHelper.InterfaceToString(value)
	if err != nil {
		return "", errors.New("PhoneNumberMustBeString")
	}

	stringValue = strings.TrimSpace(stringValue)

	re := regexp.MustCompile(phoneNumberRegex)
	isValid := re.MatchString(stringValue)
	if !isValid {
		return "", errors.New("InvalidPhoneNumber")
	}

	return PhoneNumber(stringValue), nil
}

func (number PhoneNumber) String() string {
	return string(number)
}
