package utils

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"server/internal/shared/types"

	"github.com/labstack/echo/v5"
)

func GetUser(c *echo.Context) (*types.User, bool) {
	u, ok := c.Get("user").(*types.User)
	return u, ok
}

func MustGetUser(c *echo.Context) *types.User {
	u, ok := c.Get("user").(*types.User)
	if !ok || u == nil {
		panic("user not found in context (RequireAuth middleware missing)")
	}
	return u
}

func ContextUser(c *echo.Context) (*types.User, bool) {
	u, ok := c.Get("user").(*types.User)
	return u, ok
}

func RequireUser(c *echo.Context) *types.User {
	u, ok := c.Get("user").(*types.User)
	if !ok || u == nil {
		panic("user not found in context (RequireAuth middleware missing)")
	}
	return u
}

func GenerateSessionToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return "lft_" + base64.URLEncoding.EncodeToString(b), nil
}

func FetchGitHubEmail(accessToken string) (string, error) {
	req, err := http.NewRequest("GET", "https://api.github.com/user/emails", nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer"+accessToken)
	req.Header.Set("Accept", "application/vnd.github+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var emails []struct {
		Email    string `json:"email"`
		Primary  bool   `json:"primary"`
		Verified bool   `json:"verified"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&emails); err != nil {
		return "", err
	}

	for _, e := range emails {
		if e.Primary && e.Verified {
			return e.Email, nil
		}
	}

	return "", fmt.Errorf("no verified primary email found")
}
