package utils

import (
	"encoding/json"
	"errors"
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
