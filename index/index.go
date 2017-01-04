package index

import (
	"fmt"

	"github.com/blevesearch/bleve"
	"os"
)


func DirExists(dirName string) bool {
    stat, err := os.Stat(filename)
    return err == nil || os.IsExist(err)
}

type Indexer struct {
    index bleve.Index
}

func NewIndex(indexPath string) (Indexer, error) {
    
}

// index file
func Index(indexPath string) {

	message := struct {
		Id   string
		From string
		Body string
	}{
		Id:   "example",
		From: "marty.schoch@gmail.com",
		Body: "bleve indexing is easy",
	}

	mapping := bleve.NewIndexMapping()
    stat, err := os.Stat(indexPath)
    if 


	index, err := bleve.New("example.bleve", mapping)
	if err != nil {
		panic(err)
	}
	index.Index(message.Id, message)

	query := bleve.NewQueryStringQuery("bleve")
	searchRequest := bleve.NewSearchRequest(query)
	searchResult, _ := index.Search(searchRequest)

	fmt.Printf("result %v", searchResult)
}
