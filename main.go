package injection

import (
	"bytes"
	"flag"
	"os"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/golang/glog"
)

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}

func Detect(input string) (bool, error) {
	n, err := goquery.NewDocumentFromReader(bytes.NewReader([]byte(input)))
	if err != nil {
		glog.Error(err)
		return false, err
	}
	glog.Infof("n.Nodes %+v", n.Nodes)
	return len(n.Nodes) > 0, nil
}
