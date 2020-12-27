package searcher

import (
	"bytes"
	"index/suffixarray"
	"pulley.com/shakesearch/searcher/model"
	"regexp"
	"strings"
)

type Service interface {
	Search(request model.Request) (r model.Response)
}

type searchService struct {
	CompleteWorks        string
	SuffixArray          *suffixarray.Index
	SuffixArrayLowerCase *suffixarray.Index
}

func NewSearchService(data []byte) Service {
	return &searchService{
		CompleteWorks:        string(data),
		SuffixArray:          suffixarray.New(data),
		SuffixArrayLowerCase: suffixarray.New(bytes.ToLower(data)),
	}
}

func (s *searchService) Search(request model.Request) (r model.Response) {
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
		Count:   len(idxs),
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
