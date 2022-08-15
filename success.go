package response

import (
	"time"
)

type resData struct {
	total   int
	perPage int
	result  []map[string]interface{}
}

type response struct {
	message       string
	error         string
	version       string
	representedAt string
	data          resData
	httpCode      int
}

func NewSuccessResponse() *response {
	return &response{
		version:       "v1",
		representedAt: time.Now().Format("2006-01-02 15:04:05.000000"),
	}
}

func (r *response) Message(msg string) *response {
	r.message = msg
	return r
}

func (r *response) Version(v string) *response {
	r.version = v
	return r
}

func (r *response) Data(d []map[string]interface{}) *response {
	for _, data := range d {
		r.data.result = append(r.data.result, data)
	}
	return r
}

func (r *response) SingleData(d map[string]interface{}) *response {
	r.data.result = append(r.data.result, d)
	return r
}

func (r *response) Total(t int) *response {
	r.data.total = t
	return r
}

func (r *response) Error(e string) *response {
	r.error = e
	return r
}

func (r *response) PerPage(pp int) *response {
	r.data.perPage = pp
	return r
}

func (r *response) HttpCode(hc int) *response {
	r.httpCode = hc
	return r
}

func (r *response) Generate() map[string]interface{} {
	if len(r.data.result) == 0 {
		r.data.result = []map[string]interface{}{}
	}
	resp := map[string]interface{}{
		"message":        r.message,
		"error":          r.error,
		"version":        r.version,
		"represented_at": r.representedAt,
		"data": map[string]interface{}{
			"total":    r.data.total,
			"per_page": r.data.perPage,
			"result":   r.data.result,
		},
	}

	return resp
}
