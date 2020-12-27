package searcher

import (
	"bytes"
	"encoding/json"
	"index/suffixarray"
	"net/http"
	"pulley.com/shakesearch/searcher/model"
	"regexp"
	"strconv"
	"strings"
)

type Searcher struct {
	CompleteWorks        string
	SuffixArray          *suffixarray.Index
	SuffixArrayLowerCase *suffixarray.Index
}

func NewSearcher(data []byte) *Searcher {
	return &Searcher{
		CompleteWorks:        string(data),
		SuffixArray:          suffixarray.New(data),
		SuffixArrayLowerCase: suffixarray.New(bytes.ToLower(data)),
	}
}

func (s *Searcher) HandleSearch() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var request model.Request

		query := r.URL.Query()

		q := query.Get("q")
		if len(q) < 1 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("missing search query in URL params"))
			return
		}
		request.Query = q

		sensitive := query.Get("sensitive")
		isCaseSensitive, err := strconv.ParseBool(sensitive)
		if err != nil {
			request.IsCaseSensitive = false
		} else {
			request.IsCaseSensitive = isCaseSensitive
		}

		result := s.search(request)
		buf := &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		err = enc.Encode(result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("encoding failure"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(buf.Bytes())
	}
}

func (s *Searcher) search(request model.Request) (r model.Response) {
	var idxs []int
	var toReplace *regexp.Regexp
	if request.IsCaseSensitive {
		toReplace = regexp.MustCompile(request.Query)
		idxs = s.SuffixArray.Lookup([]byte(request.Query), -1)
	} else {
		toReplace = regexp.MustCompile("(?i)" + request.Query)
		idxs = s.SuffixArrayLowerCase.Lookup([]byte(strings.ToLower(request.Query)), -1)
	}
	r = model.Response{
		Count: len(idxs),
		Matches: []string{},
	}
	for _, idx := range idxs {
		text := s.CompleteWorks[idx-250 : idx+250]
		highlightedText := toReplace.ReplaceAllStringFunc(text, func(s string) string {
			return "<highlight>" + s + "</highlight>"
		})
		r.Matches = append(r.Matches, highlightedText)
	}
	return
}
