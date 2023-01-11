package zsi

import (
	"bytes"
	"net/http"
	"net/url"
	"time"
)

func (z Zsi) fireReq(rp tRequestProcessor) {
	urlObj, err := url.Parse(rp.URL)
	if err != nil {
		rp.Errors = append(rp.Errors, err)
	}

	client := &http.Client{
		Timeout: 1 * time.Second,
	}
	request, err := http.NewRequest(
		rp.Method, urlObj.String(), bytes.NewBuffer(z.readFile(rp.Path)),
	)

	if err != nil {
		rp.Errors = append(rp.Errors, err)
	}

	if err == nil {
		request.Header.Set("Authorization", "Basic "+z.Conf.API.AuthToken)
		request.Header.Set("User-Agent", z.Conf.API.UA)
		request.Header.Set("Accept", "application/json")
		request.Header.Set("Content-Type", "application/json; charset=UTF-8")

		response, err := client.Do(request)
		if err != nil {
			rp.Errors = append(rp.Errors, err)
		} else {
			rp.ResponseStatus = response.StatusCode
		}
	}
	<-z.ChQueue
	z.ChDone <- rp
}
