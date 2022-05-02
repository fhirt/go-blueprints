package main

var thesaurus Thesaurus = UseSynonymsApi

func main() {
	UseSynonymsApi.GetSynonyms("test")
}
