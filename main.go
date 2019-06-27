package injection

import (
	"bytes"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
}

func Detect(input string) (bool, error) {
	n, err := goquery.NewDocumentFromReader(bytes.NewReader([]byte(input)))
	if err != nil {
		log.Printf("couldn't read from input string: %+v\n", err)
		return false, err
	}
	return n.Find("script").Length() > 0 || n.Find("meta").Length() > 0, nil
}
