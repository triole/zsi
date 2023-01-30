package zsi

import (
	"encoding/json"
	"strings"
	"zsi/src/conf"
)

func (z *Zsi) MakeDocList() {
	for _, indexer := range z.Conf.Indexers {
		files := z.find(indexer.Folder, indexer.RxMatcher)
		for _, fil := range files {
			if strings.HasSuffix(fil, ".csv") {
				rows := z.csvToMap(z.readFile(fil))
				for _, row := range rows {
					jsonBytes, err := json.Marshal(row)
					if err == nil {
						z.Documents = append(
							z.Documents,
							conf.Document{
								Data:  jsonBytes,
								ID:    z.getMD5Hash(string(jsonBytes)),
								Index: indexer.Index,
							},
						)
					}
				}
			} else {
				z.Documents = append(
					z.Documents,
					conf.Document{
						Path:  fil,
						ID:    z.getMD5Hash(fil),
						Index: indexer.Index,
					},
				)
			}
		}
	}
}
