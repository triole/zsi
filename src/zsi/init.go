package zsi

import (
	"zsi/src/conf"
	"zsi/src/logging"
)

type Zsi struct {
	Conf      conf.Conf
	Documents conf.Documents
	Lg        logging.Logging
	ChQueue   tChanProcessor
	ChDone    tChanProcessor
}

type tChanProcessor chan tRequestProcessor

type tRequestProcessor struct {
	Method         string
	URL            string
	Path           string
	ResponseStatus int
	Errors         []error
}

func Init(conf conf.Conf, lg logging.Logging) (zsi Zsi) {
	return Zsi{
		Conf:    conf,
		Lg:      lg,
		ChQueue: make(tChanProcessor, conf.Threads),
		ChDone:  make(tChanProcessor, conf.Threads),
	}
}
