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
"http://www.wikidata.org/entity/Q28114535" "Mr. White"
"http://www.wikidata.org/entity/Q28665865" "Мyka"
"http://www.wikidata.org/entity/Q28792126" "Gli"
"http://www.wikidata.org/entity/Q30600575" "Orlando"
"http://www.wikidata.org/entity/Q42442324" "Kiisu Miisu"
"http://www.wikidata.org/entity/Q43260736" "Paddles"
```

## Output

For each matching record, a line of output if produced, terminated with a newline character. Each field is written as a string, quoted with double quotes. Fields are separated with spaces. Double quotes and backslashes in field strings are escaped with a leading backslash.

The output can be read in Go with something like:

```
for {
   var a, b, c string
   _, err := fmt.Fscanf(stream, "%q %q %q\n", &a, &b, &c)
   if err == io.EOF {
      break
   }
   if err != nil {
      return err
   }
   ...
}
```

## Installation

Can be built with "go build" within the source directory, using a version of Go with module support (1.11 or later). Older versions may differ.
