package links

import (
	"goNiki/CheckLink/internal/dto"
	"goNiki/CheckLink/internal/infrastructure/logger/sl"
	"net/http"

	"github.com/go-chi/render"
)

func (h *handler) CheckLink(w http.ResponseWriter, r *http.Request) {

	var req dto.Req

	if err := render.DecodeJSON(r.Body, &req); err != nil {
		h.log.Error("ErrorDecode", sl.Error(err))
		render.Status(r, http.StatusBadGateway)
		return
	}

	if len(req.Links) == 0 {
		h.log.Error("links not found")
		render.Status(r, http.StatusBadRequest)
		return
	}

	links := make(map[string]string, len(req.Links))

	for _, v := range req.Links {
		status, _ := h.linkchecker.LinkCheck(r.Context(), v)
		links[v] = status
	}

	response := dto.Response{
		Links:    links,
		LinksNum: int64(len(req.Links)),
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, response)

}
