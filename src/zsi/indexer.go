package zsi

import (
	"strings"
	"zsi/src/conf"
)

func (z *Zsi) Index() {
	for _, document := range z.Documents {
		z.updateDocument(document)
	}
}

func (z Zsi) updateDocument(doc conf.Document) {
	z.fireReq(z.makeDistinctURL(z.Conf.API.UpdateDocument, doc))
}

func (z Zsi) makeDistinctURL(endpoint conf.Endpoint, document conf.Document) (method, url, path string) {
	method = strings.ToUpper(endpoint.Method)
	url = z.Conf.API.URL + document.Index + "/_doc/" + document.ID
	path = document.Path
	return
}
