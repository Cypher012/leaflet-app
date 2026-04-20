package utils

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

// GenerateUsername takes a full name and returns a slug-style username
// with a random 4-digit suffix.
func GenerateUsername(fullname string) string {
	// 1. Lowercase and trim whitespace
	username := strings.ToLower(strings.TrimSpace(fullname))

	// 2. Remove special characters (keep only letters and nsumbers)
	reg := regexp.MustCompile("[^a-z0-9]+")
	username = reg.ReplaceAllString(username, "")

	// 3. Seed random (Note: Go 1.20+ seeds automatically, but this ensures randomness)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 4. Append a random 4-digit number
	suffix := r.Intn(9000) + 1000 // Range 1000-9999
	return fmt.Sprintf("%s%d", username, suffix)
}
