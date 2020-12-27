package searcher

import (
	"bytes"
	"encoding/json"
	"index/suffixarray"
	"net/http"
	"strings"
)

type Searcher struct {
	CompleteWorks string
	SuffixArray   *suffixarray.Index
}

type Response struct {
	Count   int      `json:"count"`
	Matches []string `json:"matches"`
}

func NewSearcher(data []byte) *Searcher {
	return &Searcher{
		CompleteWorks: string(data),
		SuffixArray:   suffixarray.New(data),
	}
}

func (s *Searcher) HandleSearch() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		query, ok := r.URL.Query()["q"]
		if !ok || len(query[0]) < 1 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("missing search query in URL params"))
			return
		}
		result := s.search(query[0])
		buf := &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		err := enc.Encode(result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("encoding failure"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(buf.Bytes())
	}
}

func (s *Searcher) search(query string) (r Response) {
	idxs := s.SuffixArray.Lookup([]byte(query), -1)
	r = Response{
		Count: len(idxs),
	}
	for _, idx := range idxs {
		text := s.CompleteWorks[idx-250 : idx+250]
		highlightedText := strings.ReplaceAll(text, query, "<highlight>"+query+"</highlight>")
		r.Matches = append(r.Matches, highlightedText)
	}
	return
}
