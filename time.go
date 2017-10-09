package boson

import (
	"net/url"
	"strconv"
	"time"
)

var TimeType = struct {
	Timestamp             string
	Timedelta             string
	TimespanForTimestamps string
	TimespanForTimedeltas string
}{
	Timestamp:             "timestamp",
	Timedelta:             "timedelta",
	TimespanForTimestamps: "timespan_0",
	TimespanForTimedeltas: "timespan_1",
}

type TimeParam struct {
	Pattern  string
	BaseTime *time.Time
	noBody
}

func (t *TimeParam) query() string {
	v := url.Values{}
	v.Set("pattern", t.Pattern)
	if t.BaseTime != nil {
		v.Set("basetime", strconv.FormatInt(t.BaseTime.Unix(), 10))
	}

	return v.Encode()
}

type TimeResult struct {
	Type      string   `json:"type"`
	Timestamp string   `json:"timestamp"`
	Timedelta string   `json:"timedelta"`
	Timespan  []string `json:"timespan"`
}

type btime struct {
	b *Boson
}

func (b *btime) Analysis(param TimeParam) (*TimeResult, Response) {
	result := &TimeResult{}
	resp := b.b.do(httpMethodPost, "/time/analysis", &param, result)
	return result, resp
}
