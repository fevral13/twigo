package handlers

import (
	"github.com/fevral13/twigo/templates"
	"net/http"
)

type Index struct{}

func NewIndex() Index {
	return Index{}
}

func (i Index) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t := templates.Base(templates.Index())
	err := t.Render(r.Context(), w)
	if err != nil {

	}
}
