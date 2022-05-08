package thesaurus

type Thesaurus interface {
	GetSynonyms(string) ([]string, error)
}
