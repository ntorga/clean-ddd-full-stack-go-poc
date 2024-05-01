package valueObject

import (
	"errors"
	"regexp"
	"strings"

	voHelper "github.com/ntorga/clean-ddd-full-stack-go-poc/src/domain/valueObject/helper"
)

const nicknameRegex string = `^[\p{L}\d][\p{L}\d_-]{1,100}$`

type Nickname string

func NewNickname(value interface{}) (Nickname, error) {
	stringValue, err := voHelper.InterfaceToString(value)
	if err != nil {
		return "", errors.New("NicknameMustBeString")
	}

	stringValue = strings.TrimSpace(stringValue)
	stringValue = strings.ToLower(stringValue)

	re := regexp.MustCompile(nicknameRegex)
	isValid := re.MatchString(stringValue)
	if !isValid {
		return "", errors.New("InvalidNickname")
	}

	return Nickname(stringValue), nil
}

func (name Nickname) String() string {
	return string(name)
}
