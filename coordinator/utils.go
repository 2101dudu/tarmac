package coordinator

import (
	"crypto/rand"
	"encoding/hex"
	"strconv"
	"strings"
)

func generateToken() string {
	// defer logger.Log.TrackTime()()
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func extractPointer[T any](in *T) T {
	var out T
	if in != nil {
		out = *in
	}
	return out
}

func extractCodeAndService(in string) (string, int) {
	parts := strings.SplitN(in, "-", 2)
	codePart := parts[0]
	service, _ := strconv.Atoi(parts[1])
	return codePart, service
}
