package zsi

import (
	"fmt"
	"strings"
	"zsi/src/conf"
	"zsi/src/logging"

	"github.com/schollz/progressbar/v3"
)

func (z *Zsi) RunOperation(operation string) {
	for _, document := range z.Documents {
		if operation == "index" {
			go z.pushRequestProcessor(z.Conf.API.UpdateDocument, document)
		}
	}

	c := 0
	ll := len(z.Documents)
	if ll > 1 {
		bar := progressbar.Default(int64(ll))
		for done := range z.ChDone {
			if len(done.Errors) > 0 {
				z.Lg.Error("Request failed", logging.F{"error": done.Errors})
			}
			c++
			bar.Add(1)
			bar.Describe(z.makeBarDescription(done, len(z.ChQueue)))
			if c >= len(z.Documents) {
				close(z.ChQueue)
				close(z.ChDone)
				break
			}
		}
	} else {
		z.Lg.Info("No documents found", logging.F{})
	}
}

func (z Zsi) pushRequestProcessor(endpoint conf.Endpoint, document conf.Document) {
	rp := tRequestProcessor{
		Method:   strings.ToUpper(endpoint.Method),
		URL:      z.Conf.API.URL + document.Index + "/_doc/" + document.ID,
		Document: document,
		Errors:   []error{},
	}
	z.ChQueue <- rp
	go z.fireReq(rp)
}

func (z Zsi) makeBarDescription(rp tRequestProcessor, queueLength int) string {
	truncatedPath := z.truncateLeft(rp.Document.Path, 80)
	return fmt.Sprintf(
		"%80.80s   | t%d/%d ", truncatedPath, queueLength, z.Conf.Threads,
	)
}

func (z Zsi) truncateLeft(str string, maxLen int) string {
	if len(str) > maxLen {
		return str[len(str)-maxLen:]
	}
	return str
}
