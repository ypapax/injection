package injection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDetect(t *testing.T) {
	inputs := map[string]bool{
		`<meta http-equiv=\"refresh\" content=\"1;URL=https://some-website\" />` : true,
		`<meta http-equiv="refresh" content="1;URL=https://some-website" />` : true,
		`<script>alert('ok')</script>` : true,
		`hello world` : false,
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
