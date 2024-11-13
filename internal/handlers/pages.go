package handlers

import (
	"net/http"

	"github.com/mar-cial/food-ui/internal/templates/pages"
)

type pagesHandler struct {
}

func (h *pagesHandler) Homepage(w http.ResponseWriter, r *http.Request) {
	if err := pages.HomePage().Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func NewPagesHandler() *pagesHandler {
	return &pagesHandler{}
}
