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
	toReplace, idxs := s.lookUp(request.Query, request.IsCaseSensitive)

	r = model.Response{
		Count:   len(idxs),
		Matches: []string{},
	}

	for _, idx := range idxs {
		if request.IsExactMatch && s.isNonCompleteWord(request.Query, idx) {
			r.Count--
			continue
		}
		r.Matches = append(r.Matches, s.highlightText(idx, toReplace))
	}
	return
}

func (s *searchService) lookUp(query string, isCaseSensitive bool) (toReplace *regexp.Regexp, idxs []int) {
	if isCaseSensitive {
		toReplace = regexp.MustCompile(query)
		idxs = s.SuffixArray.Lookup([]byte(query), -1)
	} else {
		toReplace = regexp.MustCompile("(?i)" + query)
		idxs = s.SuffixArrayLowerCase.Lookup([]byte(strings.ToLower(query)), -1)
	}
	return
}

func (s *searchService) highlightText(idx int, toReplace *regexp.Regexp) string {
	startIndex := 0
	if idx-250 >= startIndex {
		startIndex = idx - 250
	}

	endIndex := len(s.CompleteWorks)
	if idx+250 <= endIndex {
		endIndex = idx + 250
	}

	text := s.CompleteWorks[startIndex:endIndex]
	highlightedText := toReplace.ReplaceAllStringFunc(text, func(s string) string {
		return "<highlight>" + s + "</highlight>"
	})
	return highlightedText
}

func (s *searchService) isNonCompleteWord(query string, idx int) bool {
	if idx-1 < 0 || idx+len(query) >= len(s.CompleteWorks) {
		return false
	}

	prefixAscii := s.CompleteWorks[idx-1]
	suffixAscii := s.CompleteWorks[idx+len(query)]

	return (isLowerCaseAlphabet(prefixAscii) || isUpperCaseAlphabet(prefixAscii)) ||
		(isLowerCaseAlphabet(suffixAscii) || isUpperCaseAlphabet(suffixAscii))
}

func isLowerCaseAlphabet(ascii uint8) bool {
	return ascii >= 97 && ascii <= 122
}

func isUpperCaseAlphabet(ascii uint8) bool {
	return ascii >= 65 && ascii <= 90
}
