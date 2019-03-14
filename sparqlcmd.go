package main

import (
	"fmt"
	"github.com/antonholmquist/jason"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func process(query string) error {
	endpoint := "https://query.wikidata.org/bigdata/namespace/wdq/sparql?"
	full := endpoint + "query=" + url.QueryEscape(query) + "&format=json"
	req, err := http.NewRequest("POST", full, nil)
	if err != nil {
		return err
	}
	req.Header.Add("User-Agent", "sparqlcmd")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	json, err := jason.NewObjectFromBytes(body)
	if err != nil {
		return err
	}
	vars, err := json.GetStringArray("head", "vars")
	if err != nil {
		return err
	}
	bindings, err := json.GetObjectArray("results", "bindings")
	if err != nil {
		return err
	}
	for i := range bindings {
		obj, err := bindings[i].Object()
		if err != nil {
			return err
		}
		sp := ""
		for j := range vars {
			val, err := obj.GetString(vars[j], "value")
			if err != nil {
				return err
			}
			fmt.Printf("%s%q", sp, val)
			sp = " "
		}
		fmt.Println()
	}
	return nil
}

func main() {
	query, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	err = process(string(query))
	if err != nil {
		panic(err)
	}
}
