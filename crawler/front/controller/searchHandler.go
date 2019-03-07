package controller

import (
	"github.com/KarenLKL/studygolang/crawler/front/view"
	"github.com/KarenLKL/studygolang/crawler/persist"
	"net/http"
	"strconv"
)

type SearchViewHandler struct {
	itemSaver *persist.ItemSaver
	view      view.SearchResultView
}

func (s SearchViewHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	q := r.FormValue("q")
	from, err := strconv.Atoi(r.FormValue("from"))
	if err != nil {
		from = 0
	}
	result, err := s.itemSaver.All("zhenai", q, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = s.view.Render(w, result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func CreateSearchViewHandler(teimplateName string, itemSaver *persist.ItemSaver) SearchViewHandler {
	return SearchViewHandler{itemSaver: itemSaver, view: view.CreateSearchResultView(teimplateName)}
}
