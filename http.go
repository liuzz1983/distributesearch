package main

import (
	"net/http"

	bleveHttp "github.com/blevesearch/bleve/http"
)

func docIDLookup(req *http.Request) string {
	return req.FormValue("docID")
}

func indexNameLookup(req *http.Request) string {
	return req.FormValue("indexName")
}

type IndexHandler struct {
	Path        string
	Method      string
	HTTPHandler http.Handler
}

func Init() []IndexHandler {
	basePath := "test"
	createIndexHandler := bleveHttp.NewCreateIndexHandler(basePath)
	createIndexHandler.IndexNameLookup = indexNameLookup

	getIndexHandler := bleveHttp.NewGetIndexHandler()
	getIndexHandler.IndexNameLookup = indexNameLookup

	deleteIndexHandler := bleveHttp.NewDeleteIndexHandler(basePath)
	deleteIndexHandler.IndexNameLookup = indexNameLookup

	return []IndexHandler{
		{
			Path:        "/list",
			Method:      "GET",
			HTTPHandler: getIndexHandler,
		},
	}
	//listIndexesHandler := bleveHttp.NewListIndexesHandler()

	/*docIndexHandler := bleveHttp.NewDocIndexHandler("")
	docIndexHandler.IndexNameLookup = indexNameLookup
	docIndexHandler.DocIDLookup = docIDLookup

	docCountHandler := bleveHttp.NewDocCountHandler("")
	docCountHandler.IndexNameLookup = indexNameLookup

	docGetHandler := bleveHttp.NewDocGetHandler("")
	docGetHandler.IndexNameLookup = indexNameLookup
	docGetHandler.DocIDLookup = docIDLookup

	docDeleteHandler := bleveHttp.NewDocDeleteHandler("")
	docDeleteHandler.IndexNameLookup = indexNameLookup
	docDeleteHandler.DocIDLookup = docIDLookup

	searchHandler := bleveHttp.NewSearchHandler("")
	searchHandler.IndexNameLookup = indexNameLookup

	listFieldsHandler := bleveHttp.NewListFieldsHandler("")
	listFieldsHandler.IndexNameLookup = indexNameLookup

	debugHandler := bleveHttp.NewDebugDocumentHandler("")
	debugHandler.IndexNameLookup = indexNameLookup
	debugHandler.DocIDLookup = docIDLookup

	aliasHandler := bleveHttp.NewAliasHandler()*/

}
