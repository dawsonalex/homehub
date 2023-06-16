package http

import (
	"io"
	"net/http"
)

type UploadState struct {
	files map[string]io.Writer
}

func chunkUploadHandler(state UploadState, writerFunc func(uploadId string) io.Writer) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.Header.Get("X-Upload-ID") == "" || r.Header.Get("Content-Range") == "" {
			w.WriteHeader(http.StatusBadRequest)
		}

	})
}
