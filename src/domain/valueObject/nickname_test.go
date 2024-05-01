package valueObject

import (
	"testing"
)

func TestNickname(t *testing.T) {
	t.Run("ValidNickname", func(t *testing.T) {
		validNicknames := []string{
			"John",
			"Jane",
			"JD",
			"JD123",
			"JD_123",
		}
		for _, name := range validNicknames {
			_, err := NewNickname(name)
			if err != nil {
				t.Errorf("[%v] ExpectedNoErrorButGotError: %v", name, err)
			}
		}
	})

	t.Run("InvalidNickname", func(t *testing.T) {
		invalidNicknames := []string{
			"",
			".",
			"..",
			"/",
			"A very long name without any reason just for the test",
			"<root>",
		}
		for _, name := range invalidNicknames {
			_, err := NewNickname(name)
			if err == nil {
				t.Errorf("[%v] ExpectedErrorButGotNoError", name)
			}
		}
	})
}
