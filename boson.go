package boson

import (
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/tidwall/gjson"
)

const (
	DefaultHost = "https://api.bosonnlp.com"

	httpMethodPost = http.MethodPost
)

type paramer interface {
	query() string
	body() (io.Reader, error)
}

type Option struct {
	Host        string
	Token       string
	HTTPTimeout time.Duration
}

func (o *Option) New() *Boson {
	if o.Host == "" {
		o.Host = DefaultHost
	}

	b := &Boson{
		httpClient: &http.Client{Timeout: o.HTTPTimeout},
		opt:        o,
	}

	b.Sentiment = sentiment{b: b}
	b.Time = btime{b: b}

	return b
}

type Boson struct {
	httpClient *http.Client
	opt        *Option

	Sentiment sentiment
	Time      btime
}

func (b *Boson) do(method string, uri string, param paramer, result interface{}) Response {
	r := Response{}

	reqBody, err := param.body()
	if err != nil {
		r.Error = err
		return r
	}

	req, err := http.NewRequest(method, b.opt.Host+uri+"?"+param.query(), reqBody)
	if err != nil {
		r.Error = err
		return r
	}

	req.Header.Set("X-Token", b.opt.Token)

	resp, err := b.httpClient.Do(req)
	if err != nil {
		r.Error = err
		return r
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		r.Error = err
		return r
	}

	r.parse(resp)

	if resp.StatusCode/100 != 2 {
		r.Error = readRequestError(data, resp.StatusCode)
		return r
	}

	r.Error = readResponseBody(data, result)
	return r
}

func readRequestError(data []byte, code int) error {
	rerr := &RequestError{
		Code:    code,
		Message: gjson.GetBytes(data, "message").String(),
	}

	if rerr.Message == "" {
		rerr.Message = http.StatusText(code)
	}

	return rerr
}

func readResponseBody(data []byte, result interface{}) error {
	if errMsg := gjson.GetBytes(data, "error"); errMsg.Type == gjson.String {
		return &MessageError{
			Err: errMsg.String(),
		}
	}

	if result != nil {
		if err := gjson.Unmarshal(data, result); err != nil {
			return err
		}

	}

	return nil
}

type noQuery struct {
}

func (n *noQuery) query() string {
	return ""
}

type noBody struct {
}

func (n *noBody) body() (io.Reader, error) {
	return nil, nil
}
