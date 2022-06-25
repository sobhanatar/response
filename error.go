package response

import (
	"time"
)

type errResponse struct {
	message       string
	error         string
	version       string
	representedAt string
	errors        []map[string]interface{}
	httpCode      int
}

func NewErrResponse() *errResponse {
	return &errResponse{
		version:       "v1",
		representedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
}

func (r *errResponse) Message(msg string) *errResponse {
	r.message = msg
	return r
}

func (r *errResponse) Version(v string) *errResponse {
	r.version = v
	return r
}

func (r *errResponse) Data(d []map[string]interface{}) *errResponse {
	for _, data := range d {
		r.errors = append(r.errors, data)
	}
	return r
}

func (r *errResponse) SingleData(d map[string]interface{}) *errResponse {
	r.errors = append(r.errors, d)
	return r
}

func (r *errResponse) Error(e string) *errResponse {
	r.error = e
	return r
}

func (r *errResponse) ValidationErrors(errors map[string][]string) *errResponse {
	for k, v := range errors {
		for _, err := range v {
			r.errors = append(r.errors, map[string]interface{}{
				"field": k,
				"error": err,
			})
		}
	}
	return r
}

func (r *errResponse) Generate() map[string]interface{} {
	resp := map[string]interface{}{
		"message":        r.message,
		"error":          r.error,
		"version":        r.version,
		"represented_at": r.representedAt,
		"data": map[string]interface{}{
			"total":    0,
			"per_page": 0,
			"result":   r.errors,
		},
	}

	return resp
}
