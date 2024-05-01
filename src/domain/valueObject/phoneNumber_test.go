package valueObject

import (
	"testing"
)

func TestPhoneNumber(t *testing.T) {
	t.Run("ValidPhoneNumber", func(t *testing.T) {
		validPhoneNumbers := []string{
			"5555555555",
			"555 555 5555",
			"555-555-5555",
			"(555) 555-5555",
			"555.555.5555",
		}
		for _, phone := range validPhoneNumbers {
			_, err := NewPhoneNumber(phone)
			if err != nil {
				t.Errorf("[%v] ExpectedNoErrorButGotError: %v", phone, err)
			}
		}
	})

	t.Run("InvalidPhoneNumber", func(t *testing.T) {
		invalidPhoneNumbers := []string{
			"",
			".",
			"..",
			"/",
			"A ph0n3 with letters",
			"<root>",
		}
		for _, phone := range invalidPhoneNumbers {
			_, err := NewPhoneNumber(phone)
			if err == nil {
				t.Errorf("[%v] ExpectedErrorButGotNoError", phone)
			}
		}
	})
}
