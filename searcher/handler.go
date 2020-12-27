package searcher

import (
	"bytes"
	"encoding/json"
	"net/http"
	"pulley.com/shakesearch/searcher/model"
	"strconv"
)

type searchHandler struct {
	searchService Service
}

func NewSearchHandler(searchService Service) searchHandler {
	return searchHandler{
		searchService: searchService,
	}
}

func (s searchHandler) HandleSearch() func(w http.ResponseWriter, r *http.Request) {
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

		result := s.searchService.Search(request)
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
