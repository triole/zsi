package zsi

import (
	"zsi/src/conf"
	"zsi/src/logging"
)

type Zsi struct {
	Conf      conf.Conf
	Documents conf.Documents
	Lg        logging.Logging
}

func Init(conf conf.Conf, lg logging.Logging) (zsi Zsi) {
	return Zsi{
		Conf: conf,
		Lg:   lg,
	}
}
