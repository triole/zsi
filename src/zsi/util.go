package zsi

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"path/filepath"
	"regexp"
	"zsi/src/logging"
)

func (z Zsi) find(basedir string, rxFilter string) []string {
	abs, err := filepath.Abs(basedir)
	z.Lg.IfErrFatal("Absolute file path creation failed", logging.F{
		"error": err,
		"path":  abs,
	})
	z.Lg.Info("Detect files", logging.F{
		"folder":    abs,
		"rxmatcher": rxFilter,
	})
	inf, err := os.Stat(abs)
	z.Lg.IfErrFatal("Can not access folder", logging.F{
		"folder": abs,
		"error":  err,
	})
	if !inf.IsDir() {
		z.Lg.IfErrFatal("Not a folder", logging.F{
			"folder": abs,
		})
	}

	filelist := []string{}
	rxf, _ := regexp.Compile(rxFilter)

	err = filepath.Walk(abs, func(path string, f os.FileInfo, err error) error {
		if rxf.MatchString(path) {
			inf, err := os.Stat(path)
			if err == nil && !inf.IsDir() {
				filelist = append(filelist, path)
			} else {
				z.Lg.IfErrInfo("Stat file failed", logging.F{
					"path": path,
				})
			}
		}
		return nil
	})
	z.Lg.IfErrFatal("Find files failed for", logging.F{
		"folder": basedir,
		"error":  err,
	})
	return filelist
}

func (z Zsi) getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func (z Zsi) readFile(path string) (content []byte) {
	abs, err := filepath.Abs(path)
	z.Conf.Lg.IfErrFatal("Absolute file path creation failed", logging.F{
		"error": err,
		"file":  abs,
	})
	content, err = os.ReadFile(abs)
	z.Conf.Lg.IfErrFatal("Can not read file", logging.F{
		"error": err,
		"file":  path,
	})
	return
}
