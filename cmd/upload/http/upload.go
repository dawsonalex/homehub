package http

import (
	"io"
	"mime"
	"net/http"
	"strconv"
	"strings"
)

type UploadState struct {
	chunks map[string][]chunk
}

type chunk struct {
	w        io.WriteCloser
	start    int64
	end      int64
	size     int64
	filename string
}

type file interface {
	io.WriteSeeker
	io.Closer
}

func chunkUploadHandler(state UploadState, writerFunc func(c chunk) (file, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Upload-ID") == "" || r.Header.Get("Content-Range") == "" {
			w.WriteHeader(http.StatusBadRequest)
		}

		/**
		 * The idea here is that this is agnostic to the storage medium. It could all be in memory,
			or all on disk, or S3. The job of this handler is to take HTTP requests and point the bytes
			in the body down the correct writer so that they can be ordered and reassembled at a later point
			in time.

			Considerations: Does this handler care about timing out unfinished uploads (e.g. after 24 hours)

		 * The general algo here is:

		 * IF there exists no partial upload for this file:
		 * 		Create the partial upload info, place to store chunks
		 * 		Set response code to HTTP 201
		 *
		 * Write this chunk to storage
		*/

		uploadId := r.Header.Get("X-Upload-ID")
		totalSize, partFrom, partTo := parseContentRange(r.Header.Get("Content-Range"))

		w.WriteHeader(http.StatusOK)

		_, params, err := mime.ParseMediaType(r.Header.Get("Content-Disposition"))
		if err != nil {
			panic(err)
		}
		filename := params["filename"]

		// TODO: decide on the chunk persistent method here
		// we can either have a generator function return writers,
		// or hardcode chunk storage on disk, but probably not both.
		thisChunk := chunk{
			size:     totalSize,
			start:    partFrom,
			end:      partTo,
			filename: filename,
		}

		f, err := writerFunc(thisChunk)
		if err != nil {
			panic(err)
		}
		_, err = f.Seek(partFrom, 0)
		if err != nil {
			panic(err)
		}

		written, err := io.CopyN(f, r.Body, totalSize)
		if err != nil {
			panic(err)
		}

	}
}

func parseContentRange(contentRange string) (totalSize int64, partFrom int64, partTo int64) {
	contentRange = strings.Replace(contentRange, "bytes ", "", -1)
	fromTo := strings.Split(contentRange, "/")[0]
	totalSize, err := strconv.ParseInt(strings.Split(contentRange, "/")[1], 10, 64)
	if err != nil {
		// TODO: handle these errors better
		panic(err)
	}

	rangeSplit := strings.Split(fromTo, "-")

	partFrom, err = strconv.ParseInt(rangeSplit[0], 10, 64)
	if err != nil {
		panic(err)
	}
	partTo, err = strconv.ParseInt(rangeSplit[1], 10, 64)
	if err != nil {
		panic(err)
	}

	return totalSize, partFrom, partTo
}
