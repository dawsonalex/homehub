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
	w        io.Writer
	start    int64
	end      int64
	size     int64
	filename string
}

func chunkUploadHandler(state UploadState, writerFunc func(uploadId string) io.Writer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Upload-ID") == "" || r.Header.Get("Content-Range") == "" {
			w.WriteHeader(http.StatusBadRequest)
		}

		/**
		 * The idea here is that this is agnostic to the storage medium. It could all be in memory,
			or all on disk, or S3. The job of this handler is to take HTTP requests and point the bytes
			in the body down the correct writer so that they can be ordered and reassembled at a later point
			in time.

			Considerations: Does this handler care about timing out unfinished uploads (e.g after 24 hours)

		 * The general algo here is:

		 * IF there exists no partial upload for this file:
		 * 		Create the partial upload info, place to store chunks
		 * 		Set response code to HTTP 201
		 *
		 * Process this chunk
		*/

		uploadId := r.Header.Get("X-Upload-ID")
		totalSize, partFrom, partTo := parseContentRange(r.Header.Get("Content-Range"))

		if _, ok := state.chunks[uploadId]; !ok {
			w.WriteHeader(http.StatusCreated)

			_, params, err := mime.ParseMediaType(r.Header.Get("Content-Disposition"))
			if err != nil {
				panic(err)
			}
			filename := params["filename"]

			// TODO: decide on the chunk persistant method here
			// we can either have a generator function return writers,
			// or hardcode chunk storage on disk, but probably not both.
			thisChunk := chunk{
				w:        writerFunc(uploadId),
				size:     totalSize,
				start:    partFrom,
				end:      partTo,
				filename: filename,
			}

			// TODO: possibly order chunks (or something) to make stitching them back together easier / faster.
			state.chunks[uploadId] = append(state.chunks[uploadId], thisChunk)
		} else {
			w.WriteHeader(http.StatusOK)
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

	splitted := strings.Split(fromTo, "-")

	partFrom, err = strconv.ParseInt(splitted[0], 10, 64)
	if err != nil {
		panic(err)
	}
	partTo, err = strconv.ParseInt(splitted[1], 10, 64)
	if err != nil {
		panic(err)
	}

	return totalSize, partFrom, partTo
}
