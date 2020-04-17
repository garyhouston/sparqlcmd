# sparqlcmd
sparqlcmd is a simple Go program that sends a SPARQL query to Wikidata and writes the result to stdout.

## Usage

It reads the query from stdin. E.g., with a file "cats" containing:

```sparql
SELECT ?item ?itemLabel 
WHERE 
{
  ?item wdt:P31 wd:Q146.
  SERVICE wikibase:label { bd:serviceParam wikibase:language "[AUTO_LANGUAGE],en". }
}
```

type:

./sparqlcmd < cats

to get output like

```
http://www.wikidata.org/entity/Q28114535,Mr. White
http://www.wikidata.org/entity/Q28665865,Myka
http://www.wikidata.org/entity/Q28792126,Gli
http://www.wikidata.org/entity/Q30600575,Orlando
http://www.wikidata.org/entity/Q42442324,Kiisu Miisu
http://www.wikidata.org/entity/Q43260736,Paddles
```

## Output

The output is in comma-separated values (CSV) format. It can be read in Go with something like:

```
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main () {
	reader := csv.NewReader(os.Stdin)
	for {
		rec, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(rec)
	}
}
```

## Installation

This package can be compiled like any simple Go application, e.g., by installing the Go compiler and in a terminal running "go build" in the source directory.
