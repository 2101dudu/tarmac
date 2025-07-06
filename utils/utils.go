package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"regexp"
)

func trimEmojis(s string) string {
	re := regexp.MustCompile(`&#\d+;`)
	return re.ReplaceAllString(s, "")
}

func trimPrefix(s string) string {
	re := regexp.MustCompile(`^\d+\s*:\s*`)
	return re.ReplaceAllString(s, "")
}

func SimplifyString(s string) string {
	return trimEmojis(trimPrefix(s))
}

func ExtractStringField(m map[string]any, key string) string {
	if val, ok := m[key].(string); ok {
		delete(m, key)
		return val
	}
	return ""
}

func GetSHA256Hash(data string) string {
	hasher := sha256.New()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}
