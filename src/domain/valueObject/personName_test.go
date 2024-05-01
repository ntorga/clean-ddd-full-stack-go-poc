package valueObject

import (
	"testing"
)

func TestPersonName(t *testing.T) {
	t.Run("ValidPersonName", func(t *testing.T) {
		validPersonNames := []string{
			"John Doe",
			"Joh'n-Doe",
			"John van Doe",
		}
		for _, name := range validPersonNames {
			_, err := NewPersonName(name)
			if err != nil {
				t.Errorf("[%v] ExpectedNoErrorButGotError: %v", name, err)
			}
		}
	})

	t.Run("InvalidPersonName", func(t *testing.T) {
		invalidPersonNames := []string{
			"",
			".",
			"..",
			"/",
			"A name with d1g1ts",
			"<root>",
		}
		for _, name := range invalidPersonNames {
			_, err := NewPersonName(name)
			if err == nil {
				t.Errorf("[%v] ExpectedErrorButGotNoError", name)
			}
		}
	})
}
