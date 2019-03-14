# sparqlcmd
sparqlcmd is a simple Go program that sends a Sparql query to Wikidata and writes the result to stdout.

## Usage

It reads the query from stdin. E.g., with a file "cats" containing:

​SELECT ?item ?itemLabel 
WHERE 
{
  ?item wdt:P31 wd:Q146.
  SERVICE wikibase:label { bd:serviceParam wikibase:language "[AUTO_LANGUAGE],en". }
}

type:
sparqlcmd < cats

to get output like:

"http://www.wikidata.org/entity/Q28114535" "Mr. White"
"http://www.wikidata.org/entity/Q28665865" "Мyka"
"http://www.wikidata.org/entity/Q28792126" "Gli"
"http://www.wikidata.org/entity/Q30600575" "Orlando"
"http://www.wikidata.org/entity/Q42442324" "Kiisu Miisu"
"http://www.wikidata.org/entity/Q43260736" "Paddles"
...

## Installation

Can be built with "go build" within the source directory.

