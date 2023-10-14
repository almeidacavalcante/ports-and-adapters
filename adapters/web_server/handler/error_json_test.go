package handler

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHandler_jsonError(t *testing.T) {
	message := "error message"
	result := jsonError(message)
	require.Equal(t, []byte(`{"message":"error message"}`), result)
}
