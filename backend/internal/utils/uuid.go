package utils

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func StringToUUID(uuidString string) (pgtype.UUID, error) {
	parsedUUID, parsingError := uuid.Parse(uuidString)
	if parsingError != nil {
		return pgtype.UUID{}, fmt.Errorf("invalid UUID: %w", parsingError)
	}

	var pgtypeUUID pgtype.UUID
	pgtypeUUID.Bytes = [16]byte(parsedUUID)
	pgtypeUUID.Valid = true

	return pgtypeUUID, nil
}

func UUIDToString(pgtypeUUID pgtype.UUID) string {
	if !pgtypeUUID.Valid {
		return ""
	}

	parsedUUID, conversionError := uuid.FromBytes(pgtypeUUID.Bytes[:])
	if conversionError != nil {
		return ""
	}

	return parsedUUID.String()
}
