package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

// private type for Context keys
type contextKey int

const (
	clientIDKey contextKey = iota
)

func credMatcher(headerName string) (mdName string, ok bool) {
	if headerName == "Login" || headerName == "Password" {
		return headerName, true
	}
	return "", false
}
