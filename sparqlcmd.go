package main

import (
	"fmt"
	"github.com/antonholmquist/jason"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func queryEntities(query string) ([][]string, error) {
	endpoint := "https://query.wikidata.org/bigdata/namespace/wdq/sparql?"
	full := endpoint + "query=" + url.QueryEscape(query) + "&format=json"
	req, err := http.NewRequest("GET", full, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", "sparqlcmd")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	json, err := jason.NewObjectFromBytes(body)
	if err != nil {
		return nil, err
	}
	vars, err := json.GetStringArray("head", "vars")
	if err != nil {
		return nil, err
	}
	bindings, err := json.GetObjectArray("results", "bindings")
	if err != nil {
		return nil, err
	}
	result := make([][]string, len(vars))
	for i := range vars {
		result[i] = make([]string, 0, 100)
	}
	for i := range bindings {
		obj, err := bindings[i].Object()
		if err != nil {
			return nil, err
		}
		for j := range vars {
			val, err := obj.GetString(vars[j], "value")
			if err != nil {
				return nil, err
			}
			result[j] = append(result[j], val)
		}
	}
	return result, nil
}

func main() {
	query, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	result, err := queryEntities(string(query))
	if err != nil {
		panic(err)
	}
	for i := range result[0] {
		sp := ""
		for j := range result {
			fmt.Printf("%s%q", sp, result[j][i])
			sp = " "
		}
		fmt.Println()
	}
}
