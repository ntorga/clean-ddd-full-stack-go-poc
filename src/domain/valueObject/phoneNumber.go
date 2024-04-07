package valueObject

import (
	"errors"
	"regexp"
)

// valid formats: (123) 1234-1234 | 123 1234-1234 | 1231234-1234 | 12312341234
const phoneNumberRegex string = `^\(?\d{1,3}\)? ?\d{1,5}-?\d{1,5}$`

type PhoneNumber string

func NewPhoneNumber(value string) (PhoneNumber, error) {
	re := regexp.MustCompile(phoneNumberRegex)
	isValid := re.MatchString(value)
	if !isValid {
		return "", errors.New("InvalidPhoneNumber")
	}

	return PhoneNumber(value), nil
}

func (number PhoneNumber) String() string {
	return string(number)
}
