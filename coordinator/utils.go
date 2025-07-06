package coordinator

import (
	"crypto/rand"
	"encoding/hex"
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
