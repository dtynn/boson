package boson

import (
	"net/http"
	"strconv"
	"time"
)

type Limit struct {
	Limit     int
	Remaining int
	Reset     time.Time
}

type Response struct {
	RequestID  string
	RateLimit  Limit
	CountLimit Limit
	Error      error
}

func (r *Response) parse(resp *http.Response) {
	r.RequestID = resp.Header.Get("X-Request-Id")

	r.RateLimit.Limit, _ = strconv.Atoi(resp.Header.Get("X-Rate-Limit-Limit"))
	r.RateLimit.Remaining, _ = strconv.Atoi(resp.Header.Get("X-Rate-Limit-Remaining"))
	rateReset, _ := strconv.ParseInt(resp.Header.Get("X-Rate-Limit-Reset"), 10, 64)
	r.RateLimit.Reset = time.Unix(rateReset, 0)

	r.CountLimit.Limit, _ = strconv.Atoi(resp.Header.Get("X-Count-Limit-Limit"))
	r.CountLimit.Remaining, _ = strconv.Atoi(resp.Header.Get("X-Count-Limit-Remaining"))
	countReset, _ := strconv.ParseInt(resp.Header.Get("X-Count-Limit-Reset"), 10, 64)
	r.CountLimit.Reset = time.Unix(countReset, 0)
}
