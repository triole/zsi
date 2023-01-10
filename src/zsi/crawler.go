package zsi

import (
	"zsi/src/conf"
)

func (z *Zsi) MakeDocList() {
	for _, indexer := range z.Conf.Indexers {
		files := z.find(indexer.Folder, indexer.RxMatcher)
		for _, fil := range files {
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
