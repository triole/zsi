package zsi

import (
	"bytes"
	"encoding/csv"
	"io"
	"log"
)

// CSVToMap takes a reader and returns an array of dictionaries, using the header row as the keys
func (z Zsi) csvToMap(b []byte) []map[string]string {
	r := csv.NewReader(bytes.NewReader(b))
	rows := []map[string]string{}
	var header []string
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if header == nil {
			header = record
		} else {
			dict := map[string]string{}
			for i := range header {
				dict[header[i]] = record[i]
			}
			rows = append(rows, dict)
		}
	}
	return rows
}
