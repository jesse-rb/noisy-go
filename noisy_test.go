package noisy

import "testing"

func TestRandomString_DefaultAlphabet(t *testing.T) {
	length := 16
	result, err := RandomString(length, nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(result) != length {
		t.Errorf("Expected length %d, got %d", length, len(result))
	}

	// Check that all characters are from the default alphabet
	defaultAlphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for _, r := range result {
		if !containsRune(defaultAlphabet, r) {
			t.Errorf("Character %q not in default alphabet", r)
		}
	}
}
