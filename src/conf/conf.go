package conf

import "zsi/src/logging"

type Conf struct {
	DB       DB        `toml:"db"`
	Indexers []Indexer `toml:"indexers"`
	Threads  int
	API      API
	Lg       logging.Logging
}

type DB struct {
	URL  string `toml:"url"`
	User string `toml:"user"`
	Pass string `toml:"pass"`
}

type Indexer struct {
	Folder    string `toml:"folder"`
	RxMatcher string `toml:"rxmatcher"`
	Index     string `toml:"index"`
}

type API struct {
	URL            string
	AuthToken      string
	UA             string
	UpdateDocument Endpoint
}

type Endpoint struct {
	URL    string
	Method string
}

type Documents []Document

type Document struct {
	ID    string
	Path  string
	Index string
}
