package valueObject

import "testing"

func TestContactId(t *testing.T) {
	t.Run("ValidContactId", func(t *testing.T) {
		validContactIds := []interface{}{
			0, 1000, 65365, "12345",
		}

		for _, id := range validContactIds {
			_, err := NewContactId(id)
			if err != nil {
				t.Errorf("[%v] ExpectedNoErrorButGotError: %v", id, err)
			}
		}
	})

	t.Run("InvalidContactId", func(t *testing.T) {
		invalidContactIds := []interface{}{
			-1, "1000X", "-455",
		}

		for _, id := range invalidContactIds {
			_, err := NewContactId(id)
			if err == nil {
				t.Errorf("[%v] ExpectedErrorButGotNoError", id)
			}
		}
	})
}
