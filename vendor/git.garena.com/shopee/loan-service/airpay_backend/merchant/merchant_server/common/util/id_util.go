package common_util

import (
	"fmt"
	"github.com/google/uuid"
)

func GenerateUUID() string {
	uuID := uuid.New()
	id, _ := uuID.MarshalBinary()
	return fmt.Sprintf("%x", id)
}
