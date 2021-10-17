package logging

import "net/http"

type Recorder struct {
	http.ResponseWriter
	StatusCode int
	Response   []byte
}

func (r *Recorder) WriteHeader(status int) {
	r.StatusCode = status
	r.ResponseWriter.WriteHeader(status)
}

func (r *Recorder) Write(data []byte) (int, error) {
	r.Response = data
	return r.ResponseWriter.Write(data)
}
