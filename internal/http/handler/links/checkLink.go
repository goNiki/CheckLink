package links

import (
	"goNiki/CheckLink/internal/dto"
	"goNiki/CheckLink/internal/http/handler/converter"
	"goNiki/CheckLink/internal/infrastructure/logger/sl"
	"net/http"

	"github.com/go-chi/render"
)

func (h *handler) CheckLink(w http.ResponseWriter, r *http.Request) {

	var req dto.ReqCheckLink

	if err := render.DecodeJSON(r.Body, &req); err != nil {
		h.log.Error("ErrorDecode", sl.Error(err))
		render.Status(r, http.StatusBadRequest)
		return
	}

	if len(req.Links) == 0 {
		h.log.Error("links not found")
		render.Status(r, http.StatusBadRequest)
		return
	}

	linkbatch, err := h.checker.CheckBatch(r.Context(), req.Links)
	if err != nil {
		h.log.Error("Error", sl.Error(err))
		render.Status(r, http.StatusInternalServerError)
		return
	}

	response := converter.LinkBatchToResponce(&linkbatch)

	render.Status(r, http.StatusOK)
	render.JSON(w, r, response)

}
