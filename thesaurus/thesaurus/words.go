package thesaurus

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Words struct {
	APIKey string
}

type synonyms struct {
	Synonyms []string `json:"synonyms"`
}

var UseSynonymsApi Words

func (w *Words) GetSynonyms(word string) ([]string, error) {

	url := "https://wordsapiv1.p.rapidapi.com/words/" + word + "/synonyms"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.New("words: Failed when building request for word api " + url + " " + err.Error())
	}

	req.Header.Add("X-RapidAPI-Host", "wordsapiv1.p.rapidapi.com")
	req.Header.Add("X-RapidAPI-Key", w.APIKey)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.New("words: Failed when looking for synonyms for " + word + " " + err.Error())
	}

	var data synonyms
	defer response.Body.Close()
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return []string{}, err
	}

	return data.Synonyms, nil
}
