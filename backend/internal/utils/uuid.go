package utils

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func StringToUUID(s string) (pgtype.UUID, error) {
	parsedUUID, err := uuid.Parse(s)
	if err != nil {
		return pgtype.UUID{}, fmt.Errorf("invalid UUID: %w", err)
	}

	var pgtypeUUID pgtype.UUID
	if err := pgtypeUUID.Scan(parsedUUID); err != nil {
		return pgtype.UUID{}, fmt.Errorf("failed to convert UUID: %w", err)
	}

	return pgtypeUUID, nil
}

func UUIDToString(u pgtype.UUID) string {
	if !u.Valid {
		return ""
	}

	parsedUUID, err := uuid.FromBytes(u.Bytes[:])
	if err != nil {
		return ""
	}

	return parsedUUID.String()
}
