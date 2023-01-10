package conf

import (
	"encoding/base64"
	"os"
	"path"
	"path/filepath"
	"zsi/src/logging"
)

func (conf *Conf) toBase64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func (conf *Conf) executablePath() string {
	e, err := os.Executable()
	conf.Lg.IfErrFatal("Can not detect binary path", logging.F{"error": err})
	return path.Dir(e)
}

func (conf *Conf) abs(path string) string {
	abs, err := filepath.Abs(path)
	conf.Lg.IfErrFatal("Absolute file path creation failed", logging.F{
		"error": err,
		"file":  path,
	})
	return abs
}

func (conf *Conf) exists(path string) bool {
	_, error := os.Open(path)
	return error == nil
}
