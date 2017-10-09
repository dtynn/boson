package boson

import (
	"testing"
)

func TestSentimentAnalysis(t *testing.T) {
	b := testClient()

	cases := []struct {
		model string
		texts []string
	}{
		{
			model: SentimentModel.Food,
			texts: []string{"这家味道还不错", "菜品太少了而且还不新鲜"},
		},
	}

	for i, c := range cases {
		p := SentimentParam{
			Model: c.model,
			Texts: c.texts,
		}

		result, resp := b.Sentiment.Analysis(p)
		if testHandleRespError(t, i, resp) {
			continue
		}

		if len(result) != len(c.texts) {
			t.Errorf("%#d expected %d pairs, got %d", i, len(c.texts), len(result))
		}

		for j, pair := range result {
			if total := pair[0] + pair[1]; total != 1 {
				t.Errorf("#%d-%d expected total 1, got %f", i+1, j+1, total)
			}
		}
	}
}
