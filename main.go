package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"pulley.com/shakesearch/searcher"
)

func main() {
	content, err := loadFile("completeworks.txt")
	if err != nil {
		log.Fatal(err)
	}

	searchService := searcher.NewSearchService(content)
	searchHandler := searcher.NewSearchHandler(searchService)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/search", searchHandler.HandleSearch())

	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	fmt.Printf("Listening on port %s...", port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func loadFile(filename string) (content []byte, err error) {
	content, err = ioutil.ReadFile(filename)
	if err != nil {
		return content, fmt.Errorf("Load: %w", err)
	}
	return
}
