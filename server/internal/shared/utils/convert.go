package utils

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func ConvertIdToPgUUID(id *string) (pgtype.UUID, error) {
	if id == nil || *id == "" {
		return pgtype.UUID{Valid: false}, nil
	}
	var uuid pgtype.UUID
	err := uuid.Scan(*id)
	return uuid, err
}

// utils/uuid.go
func PgUUIDToString(id pgtype.UUID) (string, error) {
	uid, err := uuid.FromBytes(id.Bytes[:])
	if err != nil {
		return "", err
	}
	return uid.String(), nil
}
