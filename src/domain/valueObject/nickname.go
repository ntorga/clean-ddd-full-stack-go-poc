package valueObject

import (
	"errors"
	"regexp"
)

const nicknameRegex string = `^[\p{L}_-]{2,100}$`

type Nickname string

func NewNickname(value string) (Nickname, error) {
	re := regexp.MustCompile(nicknameRegex)
	isValid := re.MatchString(value)
	if !isValid {
		return "", errors.New("InvalidNickname")
	}

	return Nickname(value), nil
}

func (name Nickname) String() string {
	return string(name)
}
