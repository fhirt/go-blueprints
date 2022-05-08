package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/fhirt/go-blueprints/thesaurus/thesaurus"
)

func main() {
	apiKey := os.Getenv("WORDS_API_KEY")
	if len(apiKey) == 0 {
		log.Fatalln("no api key set")
	}
	thesaurus := &thesaurus.Words{APIKey: apiKey}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		word := scanner.Text()
		synonyms, err := thesaurus.GetSynonyms(word)
		if err != nil {
			log.Fatalln("Failed when looking for synonyms for "+word, err)
		}
		if len(synonyms) == 0 {
			log.Fatalln("Couldn't find any synonyms for " + word)
		}
		for _, synonym := range synonyms {
			fmt.Println(synonym)
		}
	}
}
