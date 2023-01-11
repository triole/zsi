package main

import (
	"zsi/src/conf"
	"zsi/src/logging"
	"zsi/src/zsi"
)

func main() {
	parseArgs()

	lg := logging.Init(CLI.LogLevel, CLI.LogFile, CLI.LogNoColors, CLI.LogJSON)
	conf := conf.Init(CLI.ConfigFile, CLI.Threads, lg)

	zsi := zsi.Init(conf, lg)
	zsi.MakeDocList()
	zsi.RunOperation("index")
}
