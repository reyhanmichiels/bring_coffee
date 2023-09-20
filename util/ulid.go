package util

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

func GenerateUlid() (ulid.ULID, error) {
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	ulid, err := ulid.New(ms, entropy)

	if err != nil {
		return ulid, err
	}

	return ulid, nil
}
