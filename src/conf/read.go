package conf

import (
	"fmt"
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
	abs, err := filepath.Abs(filePath)
	conf.Lg.IfErrFatal("Absolute file path creation failed", logging.F{
		"error": err,
		"file":  filePath,
	})
	userHomeDir, _ := os.UserHomeDir()
	arr := []string{
		abs,
		abs + ".toml",
		path.Join(conf.executablePath()+"path", filePath),
		path.Join(userHomeDir, ".config", "zsi", filePath),
		path.Join(userHomeDir, ".conf", "zsi", filePath),
	}
	for _, fil := range arr {
		fmt.Printf("%q\n", fil)
		if conf.exists(fil) {
			s = fil
			break
		}
	}
	return
}

func (conf *Conf) readTomlFile(filename string) {
	conf.Lg.Info("Read config", logging.F{
		"file": filename,
	})
	content, err := os.ReadFile(filename)
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
