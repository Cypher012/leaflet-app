package upload

import "server/internal/shared/response"

type PresignResponse struct {
	UploadURL string `json:"upload_url"`
	PublicURL string `json:"public_url"`
}

type DocPresignResponse response.Response[PresignResponse]

var AllowedTypes = map[string]bool{
	"avatar": true,
	"feed":   true,
}
