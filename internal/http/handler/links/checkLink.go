package links

import (
	"fmt"
	"goNiki/CheckLink/internal/dto"
	"goNiki/CheckLink/internal/infrastructure/logger/sl"
	"net/http"

	"github.com/go-chi/render"
)

func (h *handler) CheckLink(w http.ResponseWriter, r *http.Request) {

	var req dto.Req

	if err := render.DecodeJSON(r.Body, &req); err != nil {
		h.log.Error("ErrorDecode", sl.Error(err))
		return
	}

	links := make(map[string]string, len(req.Links))

	for k, v := range req.Links {
		fmt.Printf("%v: %s\n", k, v)
		links[v] = "available"
	}

	response := dto.Response{
		Links:    links,
		LinksNum: int64(len(req.Links)),
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, response)

}
