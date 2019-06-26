package injection

import (
	"bytes"

	"github.com/PuerkitoBio/goquery"
	"github.com/golang/glog"
)

func Detect(input string) (bool, error) {
	n, err := goquery.NewDocumentFromReader(bytes.NewReader([]byte(input)))
	if err != nil {
		glog.Error(err)
		return false, err
	}
	return n.Find("script").Length() > 0 || n.Find("meta").Length() > 0, nil
	/*glog.Infof("n.Nodes %+v", len(n.Nodes))
	return len(n.Nodes) > 0, nil*/
}
