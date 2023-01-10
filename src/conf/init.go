package conf

import (
	"zsi/src/logging"
)

func Init(configFile string, lg logging.Logging) (conf Conf) {
	conf.Lg = lg
	configFile = conf.detectConfFile(configFile)
	conf.readTomlFile(configFile)
	return
}

func (conf *Conf) initAPI() {
	conf.API.UpdateDocument = conf.initEndpoint("_doc", "put")
}

func (conf *Conf) initEndpoint(url, method string) (ep Endpoint) {
	return Endpoint{
		URL:    conf.DB.URL + "/api/{INDEX}/{ID}",
		Method: method,
	}
}
