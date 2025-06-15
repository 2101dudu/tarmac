package render

import (
	"net/http"

	goccyjson "github.com/goccy/go-json"
)

type JSON struct {
	Data any
}

func (r JSON) Render(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	return goccyjson.NewEncoder(w).Encode(r.Data)
}

func (r JSON) WriteContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
