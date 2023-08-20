//go:build tools
// +build tools

package tools

// Tools not imported in the module but still required should be
// added here so that they are still tracked inside go.mod

import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "golang.org/x/tools/cmd/goimports"
)
