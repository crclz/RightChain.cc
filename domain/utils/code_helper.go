package utils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

var ErrNotImplemented = errors.New("NotImplemented")

func AnyAssert(t *testing.T) *require.Assertions {
	return require.New(t)
}
