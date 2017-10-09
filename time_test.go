package boson

import (
	"testing"
)

func TestTimeAnalysis(t *testing.T) {
	b := testClient()

	cases := []struct {
		pattern string
		typ     string
	}{
		{
			pattern: "2013年二月二十八日下午四点",
			typ:     TimeType.Timestamp,
		},
		{
			pattern: "2013年二月二十八日下午四点到下午九点",
			typ:     TimeType.TimespanForTimestamps,
		},
		{
			pattern: "五个小时",
			typ:     TimeType.Timedelta,
		},
		{
			pattern: "五到十个小时",
			typ:     TimeType.TimespanForTimedeltas,
		},
	}

	for i, c := range cases {
		param := TimeParam{
			Pattern: c.pattern,
		}

		res, resp := b.Time.Analysis(param)
		if testHandleRespError(t, i, resp) {
			continue
		}

		if res.Type != c.typ {
			t.Errorf("#%d expected %s, got %s", i, c.typ, res.Type)
		}
	}
}
