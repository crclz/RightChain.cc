package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

var ErrNotImplemented = errors.New("NotImplemented")

func AnyAssert(t *testing.T) *require.Assertions {
	return require.New(t)
}

func ToJson(x interface{}) string {
	var s, err = json.Marshal(x)
	if err != nil {
		panic(err)
	}

	return string(s)
}

func GetSHA256(x string) string {
	var a = sha256.Sum256([]byte(x))
	return hex.EncodeToString(a[:])
}

func GenerateSalt(bytes int) string {
	var data = make([]byte, bytes)
	_, err := rand.Read(data)
	if err != nil {
		panic(err)
	}

	return hex.EncodeToString(data)
}
