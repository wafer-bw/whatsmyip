package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetEnv(t *testing.T) {
	v := getEnv("WMIPTESTENV", "DEFAULT")
	require.Equal(t, "DEFAULT", v)
	os.Setenv("WMIPTESTENV", "NONDEFAULT")
	v = getEnv("WMIPTESTENV", "DEFAULT")
	require.Equal(t, "NONDEFAULT", v)
}
