package utils

import (
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func TestGenerateUsername(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantBase string
	}{
		{
			name:     "simple full name",
			input:    "John Doe",
			wantBase: "johndoe",
		},
		{
			name:     "name with special characters",
			input:    "John Doe!",
			wantBase: "johndoe",
		},
		{
			name:     "name with special characters",
			input:    "John Doe!",
			wantBase: "johndoe",
		},
		{
			name:     "name with extra spaces",
			input:    "  Jane Smith  ",
			wantBase: "janesmith",
		},
		{
			name:     "already lowercase",
			input:    "alice",
			wantBase: "alice",
		},
		{
			name:     "name with numbers",
			input:    "User 123",
			wantBase: "user123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GenerateUsername(tt.input)

			// must start with expected base
			if !strings.HasPrefix(result, tt.wantBase) {
				t.Errorf("got %q, want prefix %q", result, tt.wantBase)
			}

			// suffix must be exactly 4 digits in range 1000–9999
			suffix := strings.TrimPrefix(result, tt.wantBase)
			n, err := strconv.Atoi(suffix)
			if err != nil {
				t.Errorf("suffix %q is not a number", suffix)
			}
			if n < 1000 || n > 9999 {
				t.Errorf("suffix %d out of range 1000–9999", n)
			}

			// result must only contain alphanumeric characters
			matched, _ := regexp.MatchString(`^[a-z0-9]+$`, result)
			if !matched {
				t.Errorf("result %q contains non-alphanumeric characters", result)
			}
		})
	}
}
