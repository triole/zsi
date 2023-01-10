package conf

import (
	"os"
	"path"
	"path/filepath"
	"strings"
	"zsi/src/logging"

	"github.com/pelletier/go-toml"
)

func (conf *Conf) detectConfFile(filePath string) (s string) {
	if !strings.HasSuffix(".toml", filePath) {
		filePath += ".toml"
	}
	userHomeDir, _ := os.UserHomeDir()
	arr := []string{
		conf.abs(filePath),
		conf.abs(path.Join(conf.executablePath()+"path", filePath)),
		conf.abs(path.Join(userHomeDir, ".config", "zsi", filePath)),
		conf.abs(path.Join(userHomeDir, ".conf", "zsi", filePath)),
	}
	for _, fil := range arr {
		if conf.exists(fil) {
			s = fil
			break
		}
	}
	return
}

func (conf *Conf) readTomlFile(filename string) {
	abs, err := filepath.Abs(filename)
	conf.Lg.IfErrFatal("Absolute file path creation failed", logging.F{
		"error": err,
		"file":  filename,
	})
	conf.Lg.Info("Read config", logging.F{
		"file": abs,
	})
	content, err := os.ReadFile(abs)
	conf.Lg.IfErrFatal("Can not read file", logging.F{
		"error": err,
		"file":  filename,
	})
	c := Conf{}
	err = toml.Unmarshal([]byte(os.ExpandEnv(string(content))), &c)
	conf.Lg.IfErrFatal("Unable to decode toml", logging.F{
		"error": err,
	})
	conf.DB = c.DB
	conf.Indexers = c.Indexers

	conf.API.URL = conf.DB.URL + "/api/"
	conf.API.AuthToken = c.toBase64(conf.DB.User + ":" + conf.DB.Pass)
	conf.API.UA = "ZEI"

	conf.initAPI()
}
