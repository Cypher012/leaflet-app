package utils

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"server/internal/shared/response"
	"time"

	"github.com/labstack/echo/v5"
)

type Cursor struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

func EncodeCursor(id string, createdAt time.Time) string {
	c := Cursor{
		ID:        id,
		CreatedAt: createdAt,
	}

	b, _ := json.Marshal(c)
	return base64.StdEncoding.EncodeToString(b)
}

func DecodeCursor(cursor string) (*Cursor, error) {
	decoded, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "invalid cursor")
	}

	var c Cursor
	if err := json.Unmarshal(decoded, &c); err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "invalid cursor format")
	}

	return &c, nil
}

func ParseCursor(cursor string) (createdAt *time.Time, id *string, err error) {
	if cursor == "" {
		return nil, nil, nil
	}

	decoded, err := DecodeCursor(cursor)
	if err != nil {
		return nil, nil, err
	}

	return &decoded.CreatedAt, &decoded.ID, nil
}

func BuildNextCursor[T any](
	rows []T,
	limit int,
	get func(T) (string, time.Time),
) ([]T, response.PaginationMeta) {

	count := len(rows)

	if count > limit {
		last := rows[limit]
		id, createdAt := get(last)

		cRows := rows[:limit]
		cursor := EncodeCursor(id, createdAt)

		return cRows, response.PaginationMeta{
			NextCursor: &cursor,
			HasNext:    true,
			Count:      len(cRows),
		}
	}

	return rows, response.PaginationMeta{
		HasNext: false,
		Count:   count,
	}
}
