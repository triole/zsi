package zsi

import (
	"bytes"
	"net/http"
	"net/url"
	"zsi/src/logging"
)

func (z Zsi) fireReq(method, urlStr, path string) (data []byte) {
	urlObj, err := url.Parse(urlStr)
	z.Lg.IfErrError("Can not parse url", logging.F{"error": err})

	client := &http.Client{}
	request, err := http.NewRequest(
		method, urlObj.String(), bytes.NewBuffer(z.readFile(path)),
	)
	z.Lg.IfErrError("Init request failed", logging.F{
		"error": err,
	})

	if err == nil {
		request.Header.Set("Authorization", "Basic "+z.Conf.API.AuthToken)
		request.Header.Set("User-Agent", z.Conf.API.UA)
		request.Header.Set("Accept", "application/json")
		request.Header.Set("Content-Type", "application/json; charset=UTF-8")

		response, err := client.Do(request)
		fields := logging.F{
			"file":   path,
			"method": method,
			"url":    request.URL,
			"status": response.StatusCode,
		}
		if response.StatusCode == 200 || err != nil {
			z.Lg.Info("Request successful", fields)
		} else {
			fields["error"] = err
			z.Lg.Error("Request failed", fields)
		}
	}
	//getting the response
	// data, err = io.ReadAll(response.Body)
	// z.Lg.IfErrError("Error reading request body", logging.F{"error": err})

	return
}
