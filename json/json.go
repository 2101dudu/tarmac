package json

import (
	"bytes"
	"compress/gzip"
	"io"
	"log"

	goccyjson "github.com/goccy/go-json"
)

func Compress(data any) ([]byte, error) {
	// defer logger.Log.TrackTime()()
	var b bytes.Buffer
	w := gzip.NewWriter(&b)

	jsonData, err := goccyjson.Marshal(data)
	if err != nil {
		return nil, err
	}

	if _, err := w.Write(jsonData); err != nil {
		w.Close()
		return nil, err
	}

	if err := w.Close(); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func Uncompress[T any](data []byte) (*T, error) {
	// defer logger.Log.TrackTime()()
	if data == nil { // no data to uncompress
		log.Println("No data") // temp
		return nil, nil
	}

	r, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer r.Close()

	uncompressedData, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var result T
	err = goccyjson.Unmarshal(uncompressedData, &result)
	return &result, err
}
