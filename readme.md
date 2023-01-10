# ZSI ![example workflow](https://github.com/triole/zsi/actions/workflows/build.yaml/badge.svg)

<!--- mdtoc: toc begin -->

1. [Synopsis](#synopsis)
2. [Usage](#usage)
3. [Help](#help)<!--- mdtoc: toc end -->

## Synopsis

The Zinc Search Indexer crawls a folder for files and triggers indexing them via the HTTP API of [Zinc](https://github.com/zinclabs/zinc). There is no logic or processing. The files aren't parsed but send to the API as they are. Using JSON files seems to be most reasonable as Zinc quite nicely indexes those providing a full text search and a simple navigation on the web interface.

## Usage

Settings are applied by config files which look like the one below.

```go mdox-exec="cat examples/conf.toml"
[db]
url =  "https://zincsearch.box"
user = "username"
pass  = "super_secret_pass"

[[indexers]]
folder = "${HOME}/loads_of_json_files_about_nature"
rxmatcher = ".*/animal_.*.json$"
index = "animals"

[[indexers]]
folder = "${HOME}/loads_of_json_files_about_nature"
rxmatcher = ".*/vegetable_.*.json$"
index = "vegetables"
```

If there is a configuration file named `zsi_conf.toml` you would have to pass the path to the config file as positional argument like this:

```shell
zsi ${HOME}/zsi/zsi_conf
```

Not providing an absolute path will make ZSI look for the file automatically. These folders are used for the lookup in the following order.

```
1. inside the binary's folder
2. ${HOME}/.config/zsi
3. ${HOME}/.conf/zsi
```

Executing a command like `zsi animals` will make zsi look for a config file called `animals.toml` in the folders listed above. The first found will be processed.

## Help

```go mdox-exec="r -h"

a golang code example from spring

Arguments:
  [<config-file>]    file to process, positional arg required

Flags:
  -h, --help                      Show context-sensitive help.
  -l, --log-file="/dev/stdout"    log file
  -e, --log-level="info"          log level
  -n, --log-no-colors             disable output colours, print plain text
  -j, --log-json                  enable json log, instead of text one
  -d, --debug                     debug mode
  -V, --version-flag              display version
```
