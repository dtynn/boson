package boson

import (
	"bytes"
	"encoding/json"
	"io"
)

var SentimentModel = struct {
	General string
	Auto    string
	Kitchen string
	Food    string
	News    string
	Weibo   string
}{
	General: "",
	Auto:    "auto",
	Kitchen: "kitchen",
	Food:    "food",
	News:    "news",
	Weibo:   "weibo",
}

type SentimentParam struct {
	Model string
	Texts []string
}

func (s *SentimentParam) query() string {
	return s.Model
}

func (s *SentimentParam) body() (io.Reader, error) {
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(s.Texts); err != nil {
		return nil, err
	}

	return buf, nil
}

type SentimentResult [][2]float64

type sentiment struct {
	b *Boson
}

func (s *sentiment) Analysis(param SentimentParam) (SentimentResult, Response) {
	result := SentimentResult{}
	resp := s.b.do(httpMethodPost, "/sentiment/analysis", &param, &result)
	return result, resp
}
