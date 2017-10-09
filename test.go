package boson

import (
	"os"
	"testing"
	"time"
)

func testClient() *Boson {
	opt := Option{
		Token:       os.Getenv("BOSONNLP_TOKEN"),
		HTTPTimeout: 10 * time.Second,
	}

	return opt.New()
}

func testHandleRespError(t *testing.T, i int, resp Response) bool {
	if resp.Error != nil {
		t.Errorf("#%d <Req %s> got err: %s", i+1, resp.RequestID, resp.Error)
		return true
	}

	return false
}
