package util

import (
	"fmt"

	"github.com/google/uuid"
)

func GenerateUUID4() string {
	u, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	uuid4 := u.String()
	return uuid4
}
