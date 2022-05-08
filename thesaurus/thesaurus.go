package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Thesaurus interface {
	GetSynonyms(string) ([]string, error)
}
type SynonymsApi struct{}

var UseSynonymsApi SynonymsApi

func (s SynonymsApi) GetSynonyms(word string) ([]string, error) {
	url := "https://languagetools.p.rapidapi.com/all/%7Bword%7D"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Host", "languagetools.p.rapidapi.com")
	req.Header.Add("X-RapidAPI-Key", "2dbhX9sKEqmshhbPsqI3bNo6Nsw4p168RZhjsnh7Es1KYjkAAE")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

	return nil, nil
}
