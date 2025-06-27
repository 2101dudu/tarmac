package api

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"tarmac/logger"
)

func getSHA256Hash(data string) string {
	defer logger.Log.TrackTime()()
	hasher := sha256.New()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}

func generateToken() string {
	defer logger.Log.TrackTime()()
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func extractStringField(m map[string]any, key string) string {
	if val, ok := m[key].(string); ok {
		delete(m, key)
		return val
	}
	return ""
}

func extractPointer[T any](in *T) T {
	var out T
	if in != nil {
		out = *in
	}
	return out
}
