package injection

import (
	"flag"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}

func TestDetect(t *testing.T) {
	inputs := map[string]bool{
		`<meta http-equiv=\"refresh\" content=\"1;URL=https://some-website\" />`: true,
		`<meta http-equiv="refresh" content="1;URL=https://some-website" />`:     true,
		`<script>alert('ok')</script>`:                                           true,
		`hello world`:                                                            false,
	}
	for inp, exp := range inputs {
		t.Run(inp, func(t *testing.T) {
			as := assert.New(t)
			act, err := Detect(inp)
			if !as.NoError(err) {
				return
			}
			if !as.Equal(exp, act) {
				return
			}
		})
	}
}
