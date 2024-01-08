package validation

import "testing"

// ¿Qué tipo de test es este?
func TestIsValidEmail(t *testing.T) {
	cases := []struct {
		email string
		valid bool
	}{
		{"example@example.com", true},
		{"invalid-email", false},
		{"example@example.co.uk", true},
		{"example@.com", false},
		{"@example.com", false},
		{"", false},
	}

	for _, c := range cases {
		result := ValidateEmail(c.email)
		if result != c.valid {
			t.Errorf("isValidEmail(%q) == %v, want %v", c.email, result, c.valid)
		}
	}
}
